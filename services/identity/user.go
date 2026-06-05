// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package identity

import (
	"context"
	"errors"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/swopstar/swoptape/database"
	"github.com/swopstar/swoptape/services/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUser interface {
	ID() uint
	Reload(context.Context) error

	Username() string
	SetUsername(context.Context, string) error
	CheckPassword(string) error
	SetPassword(context.Context, string) error

	HasTOTP() bool
	CheckTOTP(string, *time.Time) error
	EnableTOTP(ctx context.Context, secret, token string, time *time.Time) error
	DisableTOTP(ctx context.Context) error

	Permissions() map[Permission]bool
	SetPermissions(context.Context, map[Permission]bool) error
	HasPermission(Permission) bool
	SetPermission(context.Context, ...Permission) error
	ClearPermission(context.Context, ...Permission) error
}

type User struct {
	service *Service
	id      uint
	data    database.User
}

func (svc *Service) newUser(data database.User) *User {
	return &User{
		service: svc,
		id:      data.ID,
		data:    data,
	}
}

func (svc *Service) CreateUser(
	ctx context.Context,
	username, password string,
	perms map[Permission]bool,
) (IUser, error) {
	if err := errors.Join(
		validateUsername(username),
		validatePassword(password),
	); err != nil {
		return nil, err
	}

	hash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	userPerms := make([]database.UserPermission, 0, len(perms))
	for perm, set := range perms {
		if set {
			userPerms = append(userPerms, database.UserPermission{
				Permission: string(perm),
			})
		}
	}

	data := database.User{
		Username:    username,
		Password:    hash,
		Permissions: userPerms,
	}

	if err := svc.userChain().Create(ctx, &data); err != nil {
		return nil, mapGORMError(err)
	}

	return svc.newUser(data), nil
}

func (svc *Service) GetUserByID(ctx context.Context, id uint) (IUser, error) {
	data, err := svc.userChain().
		Preload("Permissions", nil).
		Where("id = ?", id).
		First(ctx)
	if err != nil {
		return nil, mapGORMError(err)
	}

	return svc.newUser(data), nil
}

func (svc *Service) GetUserByUsername(ctx context.Context, username string) (IUser, error) {
	data, err := svc.userChain().
		Preload("Permissions", nil).
		Where("username = ?", username).
		First(ctx)
	if err != nil {
		return nil, mapGORMError(err)
	}

	return svc.newUser(data), nil
}

func (svc *Service) CountUsers(ctx context.Context) (uint64, error) {
	count, err := gorm.G[database.User](svc.db.Gorm).Count(ctx, "id")
	if err != nil {
		return 0, err
	}

	return uint64(count), err
}

//

func (u *User) ID() uint {
	return u.data.ID
}

func (u *User) Reload(ctx context.Context) error {
	user, err := u.chain().
		First(ctx)
	if err != nil {
		return mapGORMError(err)
	}

	u.data = user
	return nil
}

//

const (
	MinUsernameLength = 4
	MaxUsernameLength = 64
)

var ErrUsernameTooShort = errors.New("username is too short")
var ErrUsernameTooLong = errors.New("username is too long")
var ErrUsernameInvalidChars = errors.New("username contains invalid characters")
var ErrUsernameStartsWithDigit = errors.New("username cannot start with a digit")

func (u *User) Username() string {
	return u.data.Username
}

func (u *User) SetUsername(ctx context.Context, username string) error {
	if username == u.Username() {
		return nil
	}

	if err := validateUsername(username); err != nil {
		return err
	}

	return u.set(ctx, "username", username, func() { u.data.Username = username })
}

func validateUsername(username string) (err error) {
	if len(username) < MinUsernameLength {
		err = errors.Join(err, ErrUsernameTooShort)
	}

	if len(username) > MaxUsernameLength {
		err = errors.Join(err, ErrUsernameTooLong)
	}

	if len(username) > 0 && username[0] >= '0' && username[0] <= '9' {
		err = errors.Join(err, ErrUsernameStartsWithDigit)
	}

	for _, c := range username {
		if !isValidUsernameChar(c) {
			err = errors.Join(err, ErrUsernameInvalidChars)
			break
		}
	}

	return err
}

func isValidUsernameChar(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '-' || c == '_'
}

//

const (
	MinPasswordLength = 8
	MaxPasswordLength = 72
)

var (
	ErrPasswordTooShort   = errors.New("password is too short")
	ErrPasswordTooLong    = errors.New("password is too long")
	ErrPasswordNotComplex = errors.New("password must contain a lowercase letter, uppercase letter, and digit")
	ErrIncorrectPassword  = errors.New("incorrect password")
)

func (u *User) CheckPassword(password string) error {
	return checkPassword(u.data.Password, password)
}

func (u *User) SetPassword(ctx context.Context, password string) error {
	if err := validatePassword(password); err != nil {
		return err
	}

	hash, err := hashPassword(password)
	if err != nil {
		return err
	}

	return u.set(ctx, "password", hash, func() { u.data.Password = hash })
}

func validatePassword(password string) (err error) {
	if len(password) < MinPasswordLength {
		err = errors.Join(err, ErrPasswordTooShort)
	}

	if len(password) > MaxPasswordLength {
		err = errors.Join(err, ErrPasswordTooLong)
	}

	var hasLower, hasUpper, hasDigit bool
	for _, c := range password {
		switch {
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= '0' && c <= '9':
			hasDigit = true
		}
	}
	if !hasLower || !hasUpper || !hasDigit {
		err = errors.Join(err, ErrPasswordNotComplex)
	}

	return
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		return "", ErrPasswordTooLong
	} else if err != nil {
		return "", errors.Join(domain.ErrInternal, err)
	}

	return string(hash), nil
}

func checkPassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return ErrIncorrectPassword
	} else if err != nil {
		return errors.Join(domain.ErrInternal, err)
	}

	return nil
}

//

var (
	ErrTOTPInvalidPIN = errors.New("invalid TOTP PIN")
)

func (u *User) HasTOTP() bool {
	return u.data.TOTP != nil
}

func (u *User) CheckTOTP(pin string, timePtr *time.Time) error {
	if u.data.TOTP == nil {
		return nil
	}

	t := time.Now()
	if timePtr != nil {
		t = *timePtr
	}

	return validateTOTP(pin, *u.data.TOTP, t)
}

func (u *User) EnableTOTP(ctx context.Context, secret, pin string, timePtr *time.Time) error {
	t := time.Now()
	if timePtr != nil {
		t = *timePtr
	}

	if err := validateTOTP(pin, secret, t); err != nil {
		return err
	}

	return u.set(ctx, "totp", secret, func() { u.data.TOTP = &secret })
}

func (u *User) DisableTOTP(ctx context.Context) error {
	if u.data.TOTP == nil {
		return nil
	}

	return u.set(ctx, "totp", nil, func() { u.data.TOTP = nil })
}

func validateTOTP(pin, secret string, time time.Time) error {
	if valid, _ := totp.ValidateCustom(pin, secret, time, totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    6,
		Algorithm: otp.AlgorithmSHA1,
	}); !valid {
		return ErrTOTPInvalidPIN
	}

	return nil
}

//

func (u *User) Permissions() map[Permission]bool {
	p := make(map[Permission]bool, len(u.data.Permissions))

	for _, userPerm := range u.data.Permissions {
		p[Permission(userPerm.Permission)] = true
	}

	return p
}

func (u *User) SetPermissions(ctx context.Context, perms map[Permission]bool) error {
	newPerms := make([]database.UserPermission, 0, len(perms))
	for perm, set := range perms {
		if set {
			newPerms = append(newPerms, database.UserPermission{
				UserID:     u.id,
				Permission: string(perm),
			})
		}
	}

	err := u.service.db.Gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&database.UserPermission{}, "user_id = ?", u.id).Error; err != nil {
			return err
		}
		if len(newPerms) > 0 {
			return tx.Create(&newPerms).Error
		}
		return nil
	})
	if err != nil {
		return mapGORMError(err)
	}

	u.data.Permissions = newPerms
	return nil
}

func (u *User) HasPermission(perm Permission) bool {
	return u.Permissions()[perm]
}

func (u *User) SetPermission(ctx context.Context, perms ...Permission) error {
	toInsert := make([]database.UserPermission, 0, len(perms))
	for _, perm := range perms {
		if !u.HasPermission(perm) {
			toInsert = append(toInsert, database.UserPermission{
				UserID:     u.id,
				Permission: string(perm),
			})
		}
	}
	if len(toInsert) == 0 {
		return nil
	}

	if err := u.service.db.Gorm.WithContext(ctx).Create(&toInsert).Error; err != nil {
		if mapped := mapGORMError(err); errors.Is(mapped, domain.ErrUnique) {
			// Already exists — treat as no-op
			return nil
		}
		return mapGORMError(err)
	}

	u.data.Permissions = append(u.data.Permissions, toInsert...)
	return nil
}

func (u *User) ClearPermission(ctx context.Context, perms ...Permission) error {
	strs := make([]string, len(perms))
	for i, p := range perms {
		strs[i] = string(p)
	}

	err := u.service.db.Gorm.WithContext(ctx).
		Delete(&database.UserPermission{}, "user_id = ? AND permission IN ?", u.id, strs).
		Error
	if err != nil {
		return mapGORMError(err)
	}

	removing := make(map[string]bool, len(strs))
	for _, s := range strs {
		removing[s] = true
	}
	kept := u.data.Permissions[:0]
	for _, p := range u.data.Permissions {
		if !removing[p.Permission] {
			kept = append(kept, p)
		}
	}
	u.data.Permissions = kept
	return nil
}

//

func (u *User) set(ctx context.Context, col string, val any, apply func()) error {
	rowsAffected, err := u.chain().Update(ctx, col, val)
	if err != nil {
		return mapGORMError(err)
	}
	if rowsAffected != 1 {
		return domain.ErrStale
	}
	apply()
	return nil
}

func (svc *Service) userChain() gorm.Interface[database.User] {
	return gorm.G[database.User](svc.db.Gorm)
}

func (u *User) chain() gorm.ChainInterface[database.User] {
	return u.service.userChain().
		Where("id = ?", u.id)
}

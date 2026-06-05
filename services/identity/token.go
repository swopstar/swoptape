package identity

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"time"

	"github.com/swopstar/swoptape/database"
	"github.com/swopstar/swoptape/services/domain"
	"gorm.io/gorm"
)

const tokenPrefix = "swoptape_"

type IApplicationToken interface {
	ID() uint
	UserID() uint
	Name() string
	ExpiresAt() *time.Time
	LastUsedAt() *time.Time
	RevokedAt() *time.Time
	IsExpired() bool
	IsRevoked() bool
}

type ApplicationToken struct {
	service *Service
	data    database.ApplicationToken
}

func (svc *Service) newApplicationToken(data database.ApplicationToken) *ApplicationToken {
	return &ApplicationToken{service: svc, data: data}
}

func (t *ApplicationToken) ID() uint               { return t.data.ID }
func (t *ApplicationToken) UserID() uint           { return t.data.UserID }
func (t *ApplicationToken) Name() string           { return t.data.Name }
func (t *ApplicationToken) ExpiresAt() *time.Time  { return t.data.ExpiresAt }
func (t *ApplicationToken) LastUsedAt() *time.Time { return t.data.LastUsedAt }
func (t *ApplicationToken) RevokedAt() *time.Time  { return t.data.RevokedAt }

func (t *ApplicationToken) IsExpired() bool {
	if t.data.ExpiresAt == nil {
		return false
	}
	return time.Now().After(*t.data.ExpiresAt)
}

func (t *ApplicationToken) IsRevoked() bool { return t.data.RevokedAt != nil }

func generateApplicationToken() (raw, hash string, err error) {
	var buf [32]byte
	if _, err = rand.Read(buf[:]); err != nil {
		return
	}
	raw = tokenPrefix + base64.RawURLEncoding.EncodeToString(buf[:])
	sum := sha256.Sum256([]byte(raw))
	hash = hex.EncodeToString(sum[:])
	return
}

// CreateToken creates a new application token for the given user. The raw token
// is returned once and never stored — only its hash is kept.
func (svc *Service) CreateToken(ctx context.Context, userID uint, name string, expiresAt *time.Time) (IApplicationToken, string, error) {
	raw, hash, err := generateApplicationToken()
	if err != nil {
		return nil, "", errors.Join(domain.ErrInternal, err)
	}

	data := database.ApplicationToken{
		UserID:    userID,
		Name:      name,
		TokenHash: hash,
		ExpiresAt: expiresAt,
	}

	if err := gorm.G[database.ApplicationToken](svc.db.Gorm).Create(ctx, &data); err != nil {
		return nil, "", mapGORMError(err)
	}

	return svc.newApplicationToken(data), raw, nil
}

func (svc *Service) GetTokenByRaw(ctx context.Context, rawToken string) (IApplicationToken, error) {
	hash := svc.hashToken(rawToken)

	data, err := gorm.G[database.ApplicationToken](svc.db.Gorm).
		Where("token_hash = ?", hash).
		First(ctx)
	if err != nil {
		return nil, mapGORMError(err)
	}

	now := time.Now()
	_ = svc.db.Gorm.WithContext(ctx).
		Model(&database.ApplicationToken{}).
		Where("id = ?", data.ID).
		Update("last_used_at", now).Error

	data.LastUsedAt = &now
	return svc.newApplicationToken(data), nil
}

func (svc *Service) ListUserTokens(ctx context.Context, userID uint) ([]IApplicationToken, error) {
	rows, err := gorm.G[database.ApplicationToken](svc.db.Gorm).
		Where("user_id = ?", userID).
		Find(ctx)
	if err != nil {
		return nil, mapGORMError(err)
	}

	tokens := make([]IApplicationToken, len(rows))
	for i, row := range rows {
		tokens[i] = svc.newApplicationToken(row)
	}
	return tokens, nil
}

func (svc *Service) RevokeToken(ctx context.Context, tokenID uint) error {
	now := time.Now()
	result := svc.db.Gorm.WithContext(ctx).
		Model(&database.ApplicationToken{}).
		Where("id = ?", tokenID).
		Update("revoked_at", now)
	if result.Error != nil {
		return mapGORMError(result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

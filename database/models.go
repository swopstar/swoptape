package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var Models = []any{
	&User{},
	&UserPermission{},
	&Session{},
}

type User struct {
	gorm.Model

	Username string `gorm:"uniqueIndex"`

	Password string
	TOTP     *string

	Permissions []UserPermission
	Sessions    []Session
}

type UserPermission struct {
	UserID     uint       `gorm:"index"`
	Permission Permission `gorm:"index"`
}

type Permission string

const (
	AdminPermission  Permission = "admin"
	UploadPermission Permission = "upload"
	TagPermission    Permission = "tag"
)

type Session struct {
	gorm.Model

	UUID             uuid.UUID `gorm:"uniqueIndex"`
	RefreshTokenHash string    `gorm:"index"`

	UserID uint `gorm:"index"`

	CreatedAtIPAddress string
	CreatedByUserAgent string

	LastUsedAt *time.Time
	RevokedAt  *time.Time
}

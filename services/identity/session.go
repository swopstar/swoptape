// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package identity

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/swopstar/swoptape/database"
	"github.com/swopstar/swoptape/services/domain"
	"gorm.io/gorm"
)

type ISession interface {
	ID() uint
	UserID() uint
	UUID() uuid.UUID
	IPAddress() string
	UserAgent() string
	LastUsedAt() *time.Time
	RevokedAt() *time.Time
	IsRevoked() bool
}

type Session struct {
	service *Service
	data    database.Session
}

func (svc *Service) newSession(data database.Session) *Session {
	return &Session{service: svc, data: data}
}

func (s *Session) ID() uint               { return s.data.ID }
func (s *Session) UserID() uint           { return s.data.UserID }
func (s *Session) UUID() uuid.UUID        { return s.data.UUID }
func (s *Session) IPAddress() string      { return s.data.CreatedAtIPAddress }
func (s *Session) UserAgent() string      { return s.data.CreatedByUserAgent }
func (s *Session) LastUsedAt() *time.Time { return s.data.LastUsedAt }
func (s *Session) RevokedAt() *time.Time  { return s.data.RevokedAt }
func (s *Session) IsRevoked() bool        { return s.data.RevokedAt != nil }

// CreateSession creates a new session for the given user and returns the session,
// a short-lived JWT access token, and an opaque refresh token.
func (svc *Service) CreateSession(ctx context.Context, user IUser, ip, userAgent string) (ISession, string, string, error) {
	raw, hash, err := svc.generateRefreshToken()
	if err != nil {
		return nil, "", "", errors.Join(domain.ErrInternal, err)
	}

	data := database.Session{
		UUID:               uuid.New(),
		RefreshTokenHash:   hash,
		UserID:             user.ID(),
		CreatedAtIPAddress: ip,
		CreatedByUserAgent: userAgent,
	}

	if err := gorm.G[database.Session](svc.db.Gorm).Create(ctx, &data); err != nil {
		return nil, "", "", mapGORMError(err)
	}

	access, err := svc.generateAccessToken(data)
	if err != nil {
		return nil, "", "", errors.Join(domain.ErrInternal, err)
	}

	return svc.newSession(data), access, raw, nil
}

func (svc *Service) GetSessionByUUID(ctx context.Context, sessionUUID uuid.UUID) (ISession, error) {
	data, err := gorm.G[database.Session](svc.db.Gorm).
		Where("uuid = ?", sessionUUID).
		First(ctx)
	if err != nil {
		return nil, mapGORMError(err)
	}
	return svc.newSession(data), nil
}

// RefreshSession validates the raw refresh token, rotates it, and returns a new
// access token and refresh token. The old refresh token is invalidated.
func (svc *Service) RefreshSession(ctx context.Context, rawRefreshToken string) (ISession, string, string, error) {
	hash := svc.hashToken(rawRefreshToken)

	data, err := gorm.G[database.Session](svc.db.Gorm).
		Where("refresh_token_hash = ?", hash).
		First(ctx)
	if err != nil {
		return nil, "", "", mapGORMError(err)
	}

	if data.RevokedAt != nil {
		return nil, "", "", domain.ErrNotFound
	}

	ttl := svc.config.Auth.JWT.RefreshTokenTTL.Duration
	if time.Now().After(data.CreatedAt.Add(ttl)) {
		return nil, "", "", domain.ErrNotFound
	}

	newRaw, newHash, err := svc.generateRefreshToken()
	if err != nil {
		return nil, "", "", errors.Join(domain.ErrInternal, err)
	}

	now := time.Now()
	err = svc.db.Gorm.WithContext(ctx).
		Model(&database.Session{}).
		Where("uuid = ?", data.UUID).
		Updates(map[string]any{
			"refresh_token_hash": newHash,
			"last_used_at":       now,
		}).Error
	if err != nil {
		return nil, "", "", mapGORMError(err)
	}

	data.RefreshTokenHash = newHash
	data.LastUsedAt = &now

	access, err := svc.generateAccessToken(data)
	if err != nil {
		return nil, "", "", errors.Join(domain.ErrInternal, err)
	}

	return svc.newSession(data), access, newRaw, nil
}

func (svc *Service) RevokeSession(ctx context.Context, sessionUUID uuid.UUID) error {
	now := time.Now()
	result := svc.db.Gorm.WithContext(ctx).
		Model(&database.Session{}).
		Where("uuid = ?", sessionUUID).
		Update("revoked_at", now)
	if result.Error != nil {
		return mapGORMError(result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (svc *Service) ListUserSessions(ctx context.Context, userID uint) ([]ISession, error) {
	rows, err := gorm.G[database.Session](svc.db.Gorm).
		Where("user_id = ?", userID).
		Find(ctx)
	if err != nil {
		return nil, mapGORMError(err)
	}

	sessions := make([]ISession, len(rows))
	for i, row := range rows {
		sessions[i] = svc.newSession(row)
	}
	return sessions, nil
}

func (svc *Service) RevokeAllUserSessions(ctx context.Context, userID uint) error {
	now := time.Now()
	return mapGORMError(svc.db.Gorm.WithContext(ctx).
		Model(&database.Session{}).
		Where("user_id = ? AND revoked_at IS NULL", userID).
		Update("revoked_at", now).Error)
}

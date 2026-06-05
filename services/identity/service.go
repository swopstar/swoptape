// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package identity

import (
	"context"
	"crypto/ed25519"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/swopstar/swoptape/config"
	"github.com/swopstar/swoptape/database"
)

type IService interface {
	GetUserByID(context.Context, uint) (IUser, error)
	GetUserByUsername(context.Context, string) (IUser, error)
	CountUsers(context.Context) (uint64, error)

	CreateSession(ctx context.Context, user IUser, ip, userAgent string) (ISession, string, string, error)
	GetSessionByUUID(ctx context.Context, sessionUUID uuid.UUID) (ISession, error)
	RefreshSession(ctx context.Context, rawRefreshToken string) (ISession, string, string, error)
	RevokeSession(ctx context.Context, sessionUUID uuid.UUID) error
	ListUserSessions(ctx context.Context, userID uint) ([]ISession, error)
	RevokeAllUserSessions(ctx context.Context, userID uint) error
	ParseAccessToken(tokenString string) (userID uint, sessionUUID uuid.UUID, err error)

	CreateToken(ctx context.Context, userID uint, name string, expiresAt *time.Time) (IApplicationToken, string, error)
	GetTokenByRaw(ctx context.Context, rawToken string) (IApplicationToken, error)
	ListUserTokens(ctx context.Context, userID uint) ([]IApplicationToken, error)
	RevokeToken(ctx context.Context, tokenID uint) error
}

type Service struct {
	config *config.Config
	logger *slog.Logger
	db     *database.Database
	jwtKey ed25519.PrivateKey
}

const (
	defaultAdminUsername = "admin"
	defaultAdminPassword = "Admin123"
)

func New(
	ctx context.Context,
	cfg *config.Config,
	l *slog.Logger,
	db *database.Database,
) (*Service, error) {
	svc := &Service{
		config: cfg,
		logger: l,
		db:     db,
	}

	key, err := svc.loadOrCreateJWTKey()
	if err != nil {
		return nil, err
	}
	svc.jwtKey = key

	if err := svc.ensureAdminExists(ctx); err != nil {
		return nil, err
	}

	return svc, nil
}

func (svc *Service) ensureAdminExists(ctx context.Context) error {
	count, err := svc.CountUsers(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	_, err = svc.CreateUser(ctx, defaultAdminUsername, defaultAdminPassword, map[Permission]bool{
		PermAdmin: true,
	})
	return err
}

// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package identity

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	goenv "github.com/swopstar/gokit/config"
	"github.com/swopstar/gokit/fsutil"
	"github.com/swopstar/swoptape/config"
	"github.com/swopstar/swoptape/database"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (svc *Service) loadOrCreateJWTKey() (ed25519.PrivateKey, error) {
	keyPath := filepath.Join(config.DataDir(goenv.NewRealEnv()), svc.config.Auth.JWT.KeyPath)

	data, err := os.ReadFile(keyPath)
	if errors.Is(err, os.ErrNotExist) {
		return svc.generateAndSaveKey(keyPath)
	} else if err != nil {
		return nil, fmt.Errorf("identity: read JWT key: %w", err)
	}

	return parseEd25519PEM(data)
}

func (svc *Service) generateAndSaveKey(keyPath string) (ed25519.PrivateKey, error) {
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("identity: generate JWT key: %w", err)
	}

	der, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return nil, fmt.Errorf("identity: marshal JWT key: %w", err)
	}

	pemData := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der})

	if err := fsutil.EnsureFile(keyPath, pemData, 0600); err != nil {
		return nil, fmt.Errorf("identity: write JWT key: %w", err)
	}

	// Re-read the file in case it was written concurrently
	raw, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("identity: read JWT key after write: %w", err)
	}
	return parseEd25519PEM(raw)
}

func parseEd25519PEM(data []byte) (ed25519.PrivateKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("identity: JWT key: invalid PEM data")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("identity: parse JWT key: %w", err)
	}
	ed, ok := key.(ed25519.PrivateKey)
	if !ok {
		return nil, errors.New("identity: JWT key: not an Ed25519 key")
	}
	return ed, nil
}

func (svc *Service) generateAccessToken(sess database.Session) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    svc.config.Auth.JWT.Issuer,
		Subject:   strconv.FormatUint(uint64(sess.UserID), 10),
		ID:        sess.UUID.String(),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(svc.config.Auth.JWT.AccessTokenTTL.Duration)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	return token.SignedString(svc.jwtKey)
}

func (svc *Service) generateRefreshToken() (raw, hash string, err error) {
	var buf [32]byte
	if _, err = rand.Read(buf[:]); err != nil {
		return
	}
	raw = base64.RawURLEncoding.EncodeToString(buf[:])
	sum := sha256.Sum256([]byte(raw))
	hash = hex.EncodeToString(sum[:])
	return
}

func (svc *Service) hashToken(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}

func (svc *Service) ParseAccessToken(tokenString string) (userID uint, sessionUUID uuid.UUID, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("identity: unexpected JWT signing method: %v", t.Header["alg"])
		}
		return svc.jwtKey.Public().(ed25519.PublicKey), nil
	})
	if err != nil {
		return 0, uuid.Nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return 0, uuid.Nil, errors.New("identity: invalid JWT claims")
	}

	id, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return 0, uuid.Nil, fmt.Errorf("identity: JWT subject is not a valid user ID: %w", err)
	}

	sessionUUID, err = uuid.Parse(claims.ID)
	if err != nil {
		return 0, uuid.Nil, fmt.Errorf("identity: JWT ID is not a valid session UUID: %w", err)
	}

	return uint(id), sessionUUID, nil
}

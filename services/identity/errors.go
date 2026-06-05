// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package identity

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/swopstar/swoptape/services/domain"
	"gorm.io/gorm"
)

var ErrUpdateFailed = errors.New("operation")
var ErrTooShort = errors.New("too short")
var ErrTooLong = errors.New("too long")
var ErrInvalid = errors.New("invalid")

func mapGORMError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrNotFound
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return domain.ErrUnique
	}
	return errors.Join(domain.ErrInternal, err)
}

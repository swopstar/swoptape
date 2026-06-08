// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package v0

import (
	"errors"

	"github.com/swopstar/swoptape/services"
)

var errNotImplemented = errors.New("not implemented")

type Handlers struct {
	svc *services.Services
}

func NewHandlers(svc *services.Services) *Handlers {
	return &Handlers{svc: svc}
}

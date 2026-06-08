// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package v0

import "context"

func (h *Handlers) CreateSession(ctx context.Context, request CreateSessionRequestObject) (CreateSessionResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) RefreshToken(ctx context.Context, request RefreshTokenRequestObject) (RefreshTokenResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) EndCurrentSession(ctx context.Context, request EndCurrentSessionRequestObject) (EndCurrentSessionResponseObject, error) {
	return nil, errNotImplemented
}

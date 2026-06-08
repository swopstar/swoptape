// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package v0

import "context"

func (h *Handlers) ListUserSessions(ctx context.Context, request ListUserSessionsRequestObject) (ListUserSessionsResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) GetUserSession(ctx context.Context, request GetUserSessionRequestObject) (GetUserSessionResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) RevokeUserSession(ctx context.Context, request RevokeUserSessionRequestObject) (RevokeUserSessionResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) RevokeAllUserSessions(ctx context.Context, request RevokeAllUserSessionsRequestObject) (RevokeAllUserSessionsResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) ListAppPasswords(ctx context.Context, request ListAppPasswordsRequestObject) (ListAppPasswordsResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) CreateAppPassword(ctx context.Context, request CreateAppPasswordRequestObject) (CreateAppPasswordResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) RevokeAppPassword(ctx context.Context, request RevokeAppPasswordRequestObject) (RevokeAppPasswordResponseObject, error) {
	return nil, errNotImplemented
}

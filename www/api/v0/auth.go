// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package v0

import (
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/swopstar/swoptape/services/domain"
	"github.com/swopstar/swoptape/services/identity"
)

func (h *Handlers) CreateSession(ctx context.Context, request CreateSessionRequestObject) (CreateSessionResponseObject, error) {
	if request.Body == nil {
		return authError("invalid_request"), nil
	}
	body := request.Body

	user, err := h.svc.Identity.GetUserByUsername(ctx, body.Username)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return authError("incorrect_credentials"), nil
		}
		return nil, err
	}

	if err := user.CheckPassword(body.Password); err != nil {
		return authError("incorrect_credentials"), nil
	}

	// TOTP intentionally skipped — MFA not yet implemented.

	_, accessToken, refreshToken, err := h.svc.Identity.CreateSession(ctx, user, ginClientIP(ctx), "")
	if err != nil {
		return nil, err
	}

	username := user.Username()
	userID := float32(user.ID())
	granted := grantedEntitlements(user, body.Entitlements)

	return CreateSession200JSONResponse{
		AccessToken:  &accessToken,
		RefreshToken: &refreshToken,
		Entitlements: &granted,
		User: &User{
			Id:       &userID,
			Username: &username,
		},
	}, nil
}

func (h *Handlers) RefreshToken(ctx context.Context, request RefreshTokenRequestObject) (RefreshTokenResponseObject, error) {
	if request.Body == nil {
		return RefreshToken401Response{}, nil
	}

	_, accessToken, refreshToken, err := h.svc.Identity.RefreshSession(ctx, request.Body.RefreshToken)
	if err != nil {
		return RefreshToken401Response{}, nil
	}

	return RefreshToken200JSONResponse{
		AccessToken:  &accessToken,
		RefreshToken: &refreshToken,
	}, nil
}

func (h *Handlers) EndCurrentSession(ctx context.Context, _ EndCurrentSessionRequestObject) (EndCurrentSessionResponseObject, error) {
	gc, ok := ctx.(*gin.Context)
	if !ok {
		return EndCurrentSession204Response{}, nil
	}

	rawToken := strings.TrimPrefix(gc.GetHeader("Authorization"), "Bearer ")
	if rawToken == "" {
		return EndCurrentSession204Response{}, nil
	}

	_, sessionUUID, err := h.svc.Identity.ParseAccessToken(rawToken)
	if err != nil {
		return EndCurrentSession204Response{}, nil
	}

	_ = h.svc.Identity.RevokeSession(ctx, sessionUUID)
	return EndCurrentSession204Response{}, nil
}

// grantedEntitlements returns the subset of requested entitlements the user actually holds.
func grantedEntitlements(user identity.IUser, requested *[]string) []string {
	isAdmin := user.HasPermission(identity.PermAdmin)

	allowed := map[string]bool{
		"admin":    isAdmin,
		"tag":      isAdmin || user.HasPermission(identity.PermTag),
		"upload":   isAdmin || user.HasPermission(identity.PermUpload),
		"internal": isAdmin,
	}

	if requested == nil {
		granted := make([]string, 0, len(allowed))
		for ent, ok := range allowed {
			if ok {
				granted = append(granted, ent)
			}
		}
		return granted
	}

	granted := make([]string, 0, len(*requested))
	for _, ent := range *requested {
		if allowed[ent] {
			granted = append(granted, ent)
		}
	}
	return granted
}

func authError(code string) CreateSession401JSONResponse {
	var errCode Error_Error
	_ = errCode.FromErrorError0(code)
	return CreateSession401JSONResponse{Error: errCode}
}

func ginClientIP(ctx context.Context) string {
	gc, ok := ctx.(*gin.Context)
	if !ok {
		return ""
	}
	return gc.ClientIP()
}

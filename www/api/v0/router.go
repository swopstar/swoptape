// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package v0

import (
	"github.com/gin-gonic/gin"
	"github.com/swopstar/swoptape/services"
)

type APIRouter struct {
	handlers *Handlers
}

func New(svc *services.Services) *APIRouter {
	return &APIRouter{
		handlers: NewHandlers(svc),
	}
}

func (r *APIRouter) Install(g *gin.RouterGroup) {
	RegisterHandlers(g, NewStrictHandler(r.handlers, nil))
}

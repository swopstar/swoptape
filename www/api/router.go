// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swopstar/swoptape/services"
	v0 "github.com/swopstar/swoptape/www/api/v0"
)

type Router struct {
	V0 *v0.APIRouter
}

func New(svc *services.Services) *Router {
	return &Router{
		V0: v0.New(svc),
	}
}

func (r *Router) Install(g *gin.RouterGroup) {
	r.V0.Install(g.Group("/v0"))
}

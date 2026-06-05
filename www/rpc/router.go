// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package rpc

import "github.com/gin-gonic/gin"

type Router struct {
}

func New() *Router {
	return &Router{}
}

func (r *Router) Install(g *gin.RouterGroup) {}

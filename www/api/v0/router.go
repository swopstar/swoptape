// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package v0

import "github.com/gin-gonic/gin"

type APIRouter struct {
}

func New() *APIRouter {
	return &APIRouter{}
}

func (r *APIRouter) Install(g *gin.RouterGroup) {}

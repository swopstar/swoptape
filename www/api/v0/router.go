package v0

import "github.com/gin-gonic/gin"

type APIRouter struct {
}

func New() *APIRouter {
	return &APIRouter{}
}

func (r *APIRouter) Install(g *gin.RouterGroup) {}

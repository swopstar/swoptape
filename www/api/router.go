package api

import (
	"github.com/gin-gonic/gin"
	v0 "github.com/swopstar/swoptape/www/api/v0"
)

type Router struct {
	V0 *v0.APIRouter
}

func New() *Router {
	v0 := v0.New()

	return &Router{
		V0: v0,
	}
}

func (r *Router) Install(g *gin.RouterGroup) {
	r.V0.Install(g.Group("/v0"))
}

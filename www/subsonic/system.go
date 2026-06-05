package subsonic

import "github.com/gin-gonic/gin"

func (r *Router) ping(c *gin.Context)       { r.todo(c) }
func (r *Router) getLicense(c *gin.Context) { r.todo(c) }

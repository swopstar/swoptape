package subsonic

import "github.com/gin-gonic/gin"

func (r *Router) star(c *gin.Context)      { r.todo(c) }
func (r *Router) unstar(c *gin.Context)    { r.todo(c) }
func (r *Router) setRating(c *gin.Context) { r.todo(c) }
func (r *Router) scrobble(c *gin.Context)  { r.todo(c) }

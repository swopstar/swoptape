package www

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/swopstar/swoptape/frontend"
)

func (srv *WebServer) addRoutes() {
	srv.Router.GET("/_health", getHealth)

	srv.rpc.Install(srv.Router.Group("/rpc"))
	srv.api.Install(srv.Router.Group("/api"))
	srv.subsonic.Install(srv.Router.Group("/rest"))

	srv.webAppRoutes()
}

func (srv *WebServer) webAppRoutes() {
	distFS, _ := fs.Sub(frontend.Content, "dist")
	httpFS := http.FS(distFS)
	fileServer := http.FileServer(httpFS)

	srv.Router.NoRoute(func(c *gin.Context) {
		if isAPIRoute(c.Request.URL.Path) {
			c.Status(http.StatusNotFound)
			return
		}

		f, err := distFS.Open(c.Request.URL.Path[1:]) // strip leading slash
		if err == nil {
			_ = f.Close()
			fileServer.ServeHTTP(c.Writer, c.Request)
			return
		}

		c.FileFromFS("/", httpFS)
	})
}

func getHealth(c *gin.Context) {
	c.Status(200)
}

//

var apiRoutePrefixes = []string{
	"/rpc",  // frontend RPC channel
	"/api",  // our own API
	"/rest", // Subsonic API
}

func isAPIRoute(p string) bool {
	for _, routePrefix := range apiRoutePrefixes {
		if strings.HasPrefix(p, routePrefix) {
			return true
		}
	}

	return false
}

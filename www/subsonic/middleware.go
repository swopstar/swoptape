package subsonic

import "github.com/gin-gonic/gin"

type middleware struct {
}

func (m *middleware) handler(g *gin.Context) {
	// TODO: extract and use the following parameters to authenticate use
	//
	//   - u   - username
	//   - p   - auth token (plain text)
	//   - t+s - auth token (md5(password+salt))
	//   - v   - version
	//   - c   - user agent
	//   - f   - format

	g.Header("Cache-Control", "no-store")
	g.Next()
}

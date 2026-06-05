package subsonic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swopstar/swoptape/www/subsonic/response"
)

type Router struct {
	middleware *middleware
}

func New() *Router {
	return &Router{
		middleware: &middleware{},
	}
}

func (r *Router) Install(g *gin.RouterGroup) {
	g.Use(r.middleware.handler)

	// system
	g.GET("/ping", r.ping)
	g.GET("/getLicense", r.getLicense)

	// browsing
	g.GET("/getMusicFolders", r.getMusicFolders)
	g.GET("/getIndexes", r.getIndexes)
	g.GET("/getMusicDirectory", r.getMusicDirectory)
	g.GET("/getGenres", r.getGenres)
	g.GET("/getArtists", r.getArtists)
	g.GET("/getArtist", r.getArtist)
	g.GET("/getAlbum", r.getAlbum)
	g.GET("/getSong", r.getSong)
	g.GET("/getVideos", r.getVideos)
	g.GET("/getVideoInfo", r.getVideoInfo)
	g.GET("/getArtistInfo", r.getArtistInfo)
	g.GET("/getArtistInfo2", r.getArtistInfo2)
	g.GET("/getAlbumInfo", r.getAlbumInfo)
	g.GET("/getAlbumInfo2", r.getAlbumInfo2)
	g.GET("/getSimilarSongs", r.getSimilarSongs)
	g.GET("/getSimilarSongs2", r.getSimilarSongs2)
	g.GET("/getTopSongs", r.getTopSongs)

	// album/song lists
	g.GET("/getAlbumList", r.getAlbumList)
	g.GET("/getAlbumList2", r.getAlbumList2)
	g.GET("/getRandomSongs", r.getRandomSongs)
	g.GET("/getSongsByGenre", r.getSongsByGenre)
	g.GET("/getNowPlaying", r.getNowPlaying)
	g.GET("/getStarred", r.getStarred)
	g.GET("/getStarred2", r.getStarred2)

	// searching
	g.GET("/search", r.search)
	g.GET("/search2", r.search2)
	g.GET("/search3", r.search3)

	// playlists
	g.GET("/getPlaylists", r.getPlaylists)
	g.GET("/getPlaylist", r.getPlaylist)
	g.GET("/createPlaylist", r.createPlaylist)
	g.GET("/updatePlaylist", r.updatePlaylist)
	g.GET("/deletePlaylist", r.deletePlaylist)

	// media retrieval
	g.GET("/stream", r.stream)
	g.GET("/download", r.download)
	g.GET("/hls", r.hls)
	g.GET("/getCaptions", r.getCaptions)
	g.GET("/getCoverArt", r.getCoverArt)
	g.GET("/getLyrics", r.getLyrics)
	g.GET("/getAvatar", r.getAvatar)

	// media annotation
	g.GET("/star", r.star)
	g.GET("/unstar", r.unstar)
	g.GET("/setRating", r.setRating)
	g.GET("/scrobble", r.scrobble)

	// sharing
	g.GET("/getShares", r.getShares)
	g.GET("/createShare", r.createShare)
	g.GET("/updateShare", r.updateShare)
	g.GET("/deleteShare", r.deleteShare)

	// podcast
	g.GET("/getPodcasts", r.notImplemented)
	g.GET("/getNewestPodcasts", r.notImplemented)
	g.GET("/refreshPodcasts", r.notImplemented)
	g.GET("/createPodcastChannel", r.notImplemented)
	g.GET("/deletePodcastChannel", r.notImplemented)
	g.GET("/deletePodcastEpisode", r.notImplemented)
	g.GET("/downloadPodcastEpisode", r.notImplemented)

	// jukebox
	g.GET("/jukeboxControl", r.notImplemented)

	// internet radio
	g.GET("/getInternetRadioStations", r.notImplemented)
	g.GET("/createInternetRadioStation", r.notImplemented)
	g.GET("/updateInternetRadioStation", r.notImplemented)
	g.GET("/deleteInternetRadioStation", r.notImplemented)

	// chat
	g.GET("/getChatMessages", r.notImplemented)
	g.GET("/addChatMessage", r.notImplemented)

	// user management
	g.GET("/getUser", r.getUser)
	g.GET("/getUsers", r.getUsers)
	g.GET("/createUser", r.notImplemented)
	g.GET("/updateUser", r.notImplemented)
	g.GET("/deleteUser", r.notImplemented)
	g.GET("/changePassword", r.changePassword)

	// bookmarks
	g.GET("/getBookmarks", r.getBookmarks)
	g.GET("/createBookmark", r.createBookmark)
	g.GET("/deleteBookmark", r.deleteBookmark)
	g.GET("/getPlayQueue", r.getPlayQueue)
	g.GET("/savePlayQueue", r.savePlayQueue)

	// media library scanning
	g.GET("/getScanStatus", r.getScanStatus)
	g.GET("/startScan", r.startScan)
}

func (r *Router) todo(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

func (r *Router) notImplemented(c *gin.Context) {
	present(c, response.Response{
		Status: response.ResponseStatusFailed,
		Error:  response.GetError(response.ErrGeneric, "Not implemented"),
	})
}

func present(c *gin.Context, response response.Response) {
	format := c.GetString("Subsonic-Format")
	if format == "json" {
		c.JSON(http.StatusOK, response)
	} else {
		c.XML(http.StatusOK, response)
	}
}

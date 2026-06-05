// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package response

import "encoding/xml"

type ResponseStatus string

const (
	ResponseStatusOk     ResponseStatus = "ok"
	ResponseStatusFailed ResponseStatus = "failed"
)

type MediaType string

const (
	MediaTypeMusic     MediaType = "music"
	MediaTypePodcast   MediaType = "podcast"
	MediaTypeAudiobook MediaType = "audiobook"
	MediaTypeVideo     MediaType = "video"
)

type PodcastStatus string

const (
	PodcastStatusNew         PodcastStatus = "new"
	PodcastStatusDownloading PodcastStatus = "downloading"
	PodcastStatusCompleted   PodcastStatus = "completed"
	PodcastStatusError       PodcastStatus = "error"
	PodcastStatusDeleted     PodcastStatus = "deleted"
	PodcastStatusSkipped     PodcastStatus = "skipped"
)

type Response struct {
	XMLName xml.Name       `xml:"http://subsonic.org/restapi subsonic-response" json:"-"`
	Status  ResponseStatus `xml:"status,attr" json:"status"`
	Version string         `xml:"version,attr" json:"version"`

	MusicFolders          *MusicFolders          `xml:"musicFolders,omitempty" json:"musicFolders,omitempty"`
	Indexes               *Indexes               `xml:"indexes,omitempty" json:"indexes,omitempty"`
	Directory             *Directory             `xml:"directory,omitempty" json:"directory,omitempty"`
	Genres                *Genres                `xml:"genres,omitempty" json:"genres,omitempty"`
	Artists               *ArtistsID3            `xml:"artists,omitempty" json:"artists,omitempty"`
	Artist                *ArtistWithAlbumsID3   `xml:"artist,omitempty" json:"artist,omitempty"`
	Album                 *AlbumWithSongsID3     `xml:"album,omitempty" json:"album,omitempty"`
	Song                  *Child                 `xml:"song,omitempty" json:"song,omitempty"`
	Videos                *Videos                `xml:"videos,omitempty" json:"videos,omitempty"`
	VideoInfo             *VideoInfo             `xml:"videoInfo,omitempty" json:"videoInfo,omitempty"`
	NowPlaying            *NowPlaying            `xml:"nowPlaying,omitempty" json:"nowPlaying,omitempty"`
	SearchResult          *SearchResult          `xml:"searchResult,omitempty" json:"searchResult,omitempty"`
	SearchResult2         *SearchResult2         `xml:"searchResult2,omitempty" json:"searchResult2,omitempty"`
	SearchResult3         *SearchResult3         `xml:"searchResult3,omitempty" json:"searchResult3,omitempty"`
	Playlists             *Playlists             `xml:"playlists,omitempty" json:"playlists,omitempty"`
	Playlist              *PlaylistWithSongs     `xml:"playlist,omitempty" json:"playlist,omitempty"`
	JukeboxStatus         *JukeboxStatus         `xml:"jukeboxStatus,omitempty" json:"jukeboxStatus,omitempty"`
	JukeboxPlaylist       *JukeboxPlaylist       `xml:"jukeboxPlaylist,omitempty" json:"jukeboxPlaylist,omitempty"`
	License               *License               `xml:"license,omitempty" json:"license,omitempty"`
	Users                 *Users                 `xml:"users,omitempty" json:"users,omitempty"`
	User                  *User                  `xml:"user,omitempty" json:"user,omitempty"`
	ChatMessages          *ChatMessages          `xml:"chatMessages,omitempty" json:"chatMessages,omitempty"`
	AlbumList             *AlbumList             `xml:"albumList,omitempty" json:"albumList,omitempty"`
	AlbumList2            *AlbumList2            `xml:"albumList2,omitempty" json:"albumList2,omitempty"`
	RandomSongs           *Songs                 `xml:"randomSongs,omitempty" json:"randomSongs,omitempty"`
	SongsByGenre          *Songs                 `xml:"songsByGenre,omitempty" json:"songsByGenre,omitempty"`
	Lyrics                *Lyrics                `xml:"lyrics,omitempty" json:"lyrics,omitempty"`
	Podcasts              *Podcasts              `xml:"podcasts,omitempty" json:"podcasts,omitempty"`
	NewestPodcasts        *NewestPodcasts        `xml:"newestPodcasts,omitempty" json:"newestPodcasts,omitempty"`
	InternetRadioStations *InternetRadioStations `xml:"internetRadioStations,omitempty" json:"internetRadioStations,omitempty"`
	Bookmarks             *Bookmarks             `xml:"bookmarks,omitempty" json:"bookmarks,omitempty"`
	PlayQueue             *PlayQueue             `xml:"playQueue,omitempty" json:"playQueue,omitempty"`
	Shares                *Shares                `xml:"shares,omitempty" json:"shares,omitempty"`
	Starred               *Starred               `xml:"starred,omitempty" json:"starred,omitempty"`
	Starred2              *Starred2              `xml:"starred2,omitempty" json:"starred2,omitempty"`
	AlbumInfo             *AlbumInfo             `xml:"albumInfo,omitempty" json:"albumInfo,omitempty"`
	ArtistInfo            *ArtistInfo            `xml:"artistInfo,omitempty" json:"artistInfo,omitempty"`
	ArtistInfo2           *ArtistInfo2           `xml:"artistInfo2,omitempty" json:"artistInfo2,omitempty"`
	SimilarSongs          *SimilarSongs          `xml:"similarSongs,omitempty" json:"similarSongs,omitempty"`
	SimilarSongs2         *SimilarSongs2         `xml:"similarSongs2,omitempty" json:"similarSongs2,omitempty"`
	TopSongs              *TopSongs              `xml:"topSongs,omitempty" json:"topSongs,omitempty"`
	ScanStatus            *ScanStatus            `xml:"scanStatus,omitempty" json:"scanStatus,omitempty"`
	Error                 *Error                 `xml:"error,omitempty" json:"error,omitempty"`
}

type MusicFolders struct {
	MusicFolder []MusicFolder `xml:"musicFolder" json:"musicFolder,omitempty"`
}

type MusicFolder struct {
	ID   int     `xml:"id,attr" json:"id"`
	Name *string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type Indexes struct {
	Shortcut        []Artist `xml:"shortcut" json:"shortcut,omitempty"`
	Index           []Index  `xml:"index" json:"index,omitempty"`
	Child           []Child  `xml:"child" json:"child,omitempty"`
	LastModified    int64    `xml:"lastModified,attr" json:"lastModified"`
	IgnoredArticles string   `xml:"ignoredArticles,attr" json:"ignoredArticles"`
}

type Index struct {
	Artist []Artist `xml:"artist" json:"artist,omitempty"`
	Name   string   `xml:"name,attr" json:"name"`
}

type Artist struct {
	ID             string    `xml:"id,attr" json:"id"`
	Name           string    `xml:"name,attr" json:"name"`
	ArtistImageURL *string   `xml:"artistImageUrl,attr,omitempty" json:"artistImageUrl,omitempty"`
	Starred        *DateTime `xml:"starred,attr,omitempty" json:"starred,omitempty"`
	UserRating     *int      `xml:"userRating,attr,omitempty" json:"userRating,omitempty"`
	AverageRating  *float64  `xml:"averageRating,attr,omitempty" json:"averageRating,omitempty"`
}

type Genres struct {
	Genre []Genre `xml:"genre" json:"genre,omitempty"`
}

type Genre struct {
	SongCount  int    `xml:"songCount,attr" json:"songCount"`
	AlbumCount int    `xml:"albumCount,attr" json:"albumCount"`
	Value      string `xml:",chardata" json:"value,omitempty"`
}

type ArtistsID3 struct {
	Index           []IndexID3 `xml:"index" json:"index,omitempty"`
	IgnoredArticles string     `xml:"ignoredArticles,attr" json:"ignoredArticles"`
}

type IndexID3 struct {
	Artist []ArtistID3 `xml:"artist" json:"artist,omitempty"`
	Name   string      `xml:"name,attr" json:"name"`
}

type ArtistID3 struct {
	ID             string    `xml:"id,attr" json:"id"`
	Name           string    `xml:"name,attr" json:"name"`
	CoverArt       *string   `xml:"coverArt,attr,omitempty" json:"coverArt,omitempty"`
	ArtistImageURL *string   `xml:"artistImageUrl,attr,omitempty" json:"artistImageUrl,omitempty"`
	AlbumCount     int       `xml:"albumCount,attr" json:"albumCount"`
	Starred        *DateTime `xml:"starred,attr,omitempty" json:"starred,omitempty"`
}

type ArtistWithAlbumsID3 struct {
	ArtistID3
	Album []AlbumID3 `xml:"album" json:"album,omitempty"`
}

type AlbumID3 struct {
	ID        string    `xml:"id,attr" json:"id"`
	Name      string    `xml:"name,attr" json:"name"`
	Artist    *string   `xml:"artist,attr,omitempty" json:"artist,omitempty"`
	ArtistID  *string   `xml:"artistId,attr,omitempty" json:"artistId,omitempty"`
	CoverArt  *string   `xml:"coverArt,attr,omitempty" json:"coverArt,omitempty"`
	SongCount int       `xml:"songCount,attr" json:"songCount"`
	Duration  int       `xml:"duration,attr" json:"duration"`
	PlayCount *int64    `xml:"playCount,attr,omitempty" json:"playCount,omitempty"`
	Created   DateTime  `xml:"created,attr" json:"created"`
	Starred   *DateTime `xml:"starred,attr,omitempty" json:"starred,omitempty"`
	Year      *int      `xml:"year,attr,omitempty" json:"year,omitempty"`
	Genre     *string   `xml:"genre,attr,omitempty" json:"genre,omitempty"`
}

type AlbumWithSongsID3 struct {
	AlbumID3
	Song []Child `xml:"song" json:"song,omitempty"`
}

type Videos struct {
	Video []Child `xml:"video" json:"video,omitempty"`
}

type VideoInfo struct {
	Captions   []Captions        `xml:"captions" json:"captions,omitempty"`
	AudioTrack []AudioTrack      `xml:"audioTrack" json:"audioTrack,omitempty"`
	Conversion []VideoConversion `xml:"conversion" json:"conversion,omitempty"`
	ID         string            `xml:"id,attr" json:"id"`
}

type Captions struct {
	ID   string  `xml:"id,attr" json:"id"`
	Name *string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type AudioTrack struct {
	ID           string  `xml:"id,attr" json:"id"`
	Name         *string `xml:"name,attr,omitempty" json:"name,omitempty"`
	LanguageCode *string `xml:"languageCode,attr,omitempty" json:"languageCode,omitempty"`
}

type VideoConversion struct {
	ID           string `xml:"id,attr" json:"id"`
	BitRate      *int   `xml:"bitRate,attr,omitempty" json:"bitRate,omitempty"`
	AudioTrackID *int   `xml:"audioTrackId,attr,omitempty" json:"audioTrackId,omitempty"`
}

type Directory struct {
	Child         []Child   `xml:"child" json:"child,omitempty"`
	ID            string    `xml:"id,attr" json:"id"`
	Parent        *string   `xml:"parent,attr,omitempty" json:"parent,omitempty"`
	Name          string    `xml:"name,attr" json:"name"`
	Starred       *DateTime `xml:"starred,attr,omitempty" json:"starred,omitempty"`
	UserRating    *int      `xml:"userRating,attr,omitempty" json:"userRating,omitempty"`
	AverageRating *float64  `xml:"averageRating,attr,omitempty" json:"averageRating,omitempty"`
	PlayCount     *int64    `xml:"playCount,attr,omitempty" json:"playCount,omitempty"`
}

type Child struct {
	ID                    string     `xml:"id,attr" json:"id"`
	Parent                *string    `xml:"parent,attr,omitempty" json:"parent,omitempty"`
	IsDir                 bool       `xml:"isDir,attr" json:"isDir"`
	Title                 string     `xml:"title,attr" json:"title"`
	Album                 *string    `xml:"album,attr,omitempty" json:"album,omitempty"`
	Artist                *string    `xml:"artist,attr,omitempty" json:"artist,omitempty"`
	Track                 *int       `xml:"track,attr,omitempty" json:"track,omitempty"`
	Year                  *int       `xml:"year,attr,omitempty" json:"year,omitempty"`
	Genre                 *string    `xml:"genre,attr,omitempty" json:"genre,omitempty"`
	CoverArt              *string    `xml:"coverArt,attr,omitempty" json:"coverArt,omitempty"`
	Size                  *int64     `xml:"size,attr,omitempty" json:"size,omitempty"`
	ContentType           *string    `xml:"contentType,attr,omitempty" json:"contentType,omitempty"`
	Suffix                *string    `xml:"suffix,attr,omitempty" json:"suffix,omitempty"`
	TranscodedContentType *string    `xml:"transcodedContentType,attr,omitempty" json:"transcodedContentType,omitempty"`
	TranscodedSuffix      *string    `xml:"transcodedSuffix,attr,omitempty" json:"transcodedSuffix,omitempty"`
	Duration              *int       `xml:"duration,attr,omitempty" json:"duration,omitempty"`
	BitRate               *int       `xml:"bitRate,attr,omitempty" json:"bitRate,omitempty"`
	Path                  *string    `xml:"path,attr,omitempty" json:"path,omitempty"`
	IsVideo               *bool      `xml:"isVideo,attr,omitempty" json:"isVideo,omitempty"`
	UserRating            *int       `xml:"userRating,attr,omitempty" json:"userRating,omitempty"`
	AverageRating         *float64   `xml:"averageRating,attr,omitempty" json:"averageRating,omitempty"`
	PlayCount             *int64     `xml:"playCount,attr,omitempty" json:"playCount,omitempty"`
	DiscNumber            *int       `xml:"discNumber,attr,omitempty" json:"discNumber,omitempty"`
	Created               *DateTime  `xml:"created,attr,omitempty" json:"created,omitempty"`
	Starred               *DateTime  `xml:"starred,attr,omitempty" json:"starred,omitempty"`
	AlbumID               *string    `xml:"albumId,attr,omitempty" json:"albumId,omitempty"`
	ArtistID              *string    `xml:"artistId,attr,omitempty" json:"artistId,omitempty"`
	Type                  *MediaType `xml:"type,attr,omitempty" json:"type,omitempty"`
	BookmarkPosition      *int64     `xml:"bookmarkPosition,attr,omitempty" json:"bookmarkPosition,omitempty"`
	OriginalWidth         *int       `xml:"originalWidth,attr,omitempty" json:"originalWidth,omitempty"`
	OriginalHeight        *int       `xml:"originalHeight,attr,omitempty" json:"originalHeight,omitempty"`
}

type NowPlaying struct {
	Entry []NowPlayingEntry `xml:"entry" json:"entry,omitempty"`
}

type NowPlayingEntry struct {
	Child
	Username   string  `xml:"username,attr" json:"username"`
	MinutesAgo int     `xml:"minutesAgo,attr" json:"minutesAgo"`
	PlayerID   int     `xml:"playerId,attr" json:"playerId"`
	PlayerName *string `xml:"playerName,attr,omitempty" json:"playerName,omitempty"`
}

type SearchResult struct {
	Match     []Child `xml:"match" json:"match,omitempty"`
	Offset    int     `xml:"offset,attr" json:"offset"`
	TotalHits int     `xml:"totalHits,attr" json:"totalHits"`
}

type SearchResult2 struct {
	Artist []Artist `xml:"artist" json:"artist,omitempty"`
	Album  []Child  `xml:"album" json:"album,omitempty"`
	Song   []Child  `xml:"song" json:"song,omitempty"`
}

type SearchResult3 struct {
	Artist []ArtistID3 `xml:"artist" json:"artist,omitempty"`
	Album  []AlbumID3  `xml:"album" json:"album,omitempty"`
	Song   []Child     `xml:"song" json:"song,omitempty"`
}

type Playlists struct {
	Playlist []Playlist `xml:"playlist" json:"playlist,omitempty"`
}

type Playlist struct {
	AllowedUser []string `xml:"allowedUser" json:"allowedUser,omitempty"`
	ID          string   `xml:"id,attr" json:"id"`
	Name        string   `xml:"name,attr" json:"name"`
	Comment     *string  `xml:"comment,attr,omitempty" json:"comment,omitempty"`
	Owner       *string  `xml:"owner,attr,omitempty" json:"owner,omitempty"`
	Public      *bool    `xml:"public,attr,omitempty" json:"public,omitempty"`
	SongCount   int      `xml:"songCount,attr" json:"songCount"`
	Duration    int      `xml:"duration,attr" json:"duration"`
	Created     DateTime `xml:"created,attr" json:"created"`
	Changed     DateTime `xml:"changed,attr" json:"changed"`
	CoverArt    *string  `xml:"coverArt,attr,omitempty" json:"coverArt,omitempty"`
}

type PlaylistWithSongs struct {
	Playlist
	Entry []Child `xml:"entry" json:"entry,omitempty"`
}

type JukeboxStatus struct {
	CurrentIndex int     `xml:"currentIndex,attr" json:"currentIndex"`
	Playing      bool    `xml:"playing,attr" json:"playing"`
	Gain         float32 `xml:"gain,attr" json:"gain"`
	Position     *int    `xml:"position,attr,omitempty" json:"position,omitempty"`
}

type JukeboxPlaylist struct {
	JukeboxStatus
	Entry []Child `xml:"entry" json:"entry,omitempty"`
}

type ChatMessages struct {
	ChatMessage []ChatMessage `xml:"chatMessage" json:"chatMessage,omitempty"`
}

type ChatMessage struct {
	Username string `xml:"username,attr" json:"username"`
	Time     int64  `xml:"time,attr" json:"time"`
	Message  string `xml:"message,attr" json:"message"`
}

type AlbumList struct {
	Album []Child `xml:"album" json:"album,omitempty"`
}

type AlbumList2 struct {
	Album []AlbumID3 `xml:"album" json:"album,omitempty"`
}

type Songs struct {
	Song []Child `xml:"song" json:"song,omitempty"`
}

type Lyrics struct {
	Artist *string `xml:"artist,attr,omitempty" json:"artist,omitempty"`
	Title  *string `xml:"title,attr,omitempty" json:"title,omitempty"`
	Value  string  `xml:",chardata" json:"value,omitempty"`
}

type Podcasts struct {
	Channel []PodcastChannel `xml:"channel" json:"channel,omitempty"`
}

type PodcastChannel struct {
	Episode          []PodcastEpisode `xml:"episode" json:"episode,omitempty"`
	ID               string           `xml:"id,attr" json:"id"`
	URL              string           `xml:"url,attr" json:"url"`
	Title            *string          `xml:"title,attr,omitempty" json:"title,omitempty"`
	Description      *string          `xml:"description,attr,omitempty" json:"description,omitempty"`
	CoverArt         *string          `xml:"coverArt,attr,omitempty" json:"coverArt,omitempty"`
	OriginalImageURL *string          `xml:"originalImageUrl,attr,omitempty" json:"originalImageUrl,omitempty"`
	Status           PodcastStatus    `xml:"status,attr" json:"status"`
	ErrorMessage     *string          `xml:"errorMessage,attr,omitempty" json:"errorMessage,omitempty"`
}

type NewestPodcasts struct {
	Episode []PodcastEpisode `xml:"episode" json:"episode,omitempty"`
}

type PodcastEpisode struct {
	Child
	StreamID    *string       `xml:"streamId,attr,omitempty" json:"streamId,omitempty"`
	ChannelID   string        `xml:"channelId,attr" json:"channelId"`
	Description *string       `xml:"description,attr,omitempty" json:"description,omitempty"`
	Status      PodcastStatus `xml:"status,attr" json:"status"`
	PublishDate *DateTime     `xml:"publishDate,attr,omitempty" json:"publishDate,omitempty"`
}

type InternetRadioStations struct {
	InternetRadioStation []InternetRadioStation `xml:"internetRadioStation" json:"internetRadioStation,omitempty"`
}

type InternetRadioStation struct {
	ID          string  `xml:"id,attr" json:"id"`
	Name        string  `xml:"name,attr" json:"name"`
	StreamURL   string  `xml:"streamUrl,attr" json:"streamUrl"`
	HomePageURL *string `xml:"homePageUrl,attr,omitempty" json:"homePageUrl,omitempty"`
}

type Bookmarks struct {
	Bookmark []Bookmark `xml:"bookmark" json:"bookmark,omitempty"`
}

type Bookmark struct {
	Entry    Child    `xml:"entry" json:"entry"`
	Position int64    `xml:"position,attr" json:"position"`
	Username string   `xml:"username,attr" json:"username"`
	Comment  *string  `xml:"comment,attr,omitempty" json:"comment,omitempty"`
	Created  DateTime `xml:"created,attr" json:"created"`
	Changed  DateTime `xml:"changed,attr" json:"changed"`
}

type PlayQueue struct {
	Entry     []Child  `xml:"entry" json:"entry,omitempty"`
	Current   *int     `xml:"current,attr,omitempty" json:"current,omitempty"`
	Position  *int64   `xml:"position,attr,omitempty" json:"position,omitempty"`
	Username  string   `xml:"username,attr" json:"username"`
	Changed   DateTime `xml:"changed,attr" json:"changed"`
	ChangedBy string   `xml:"changedBy,attr" json:"changedBy"`
}

type Shares struct {
	Share []Share `xml:"share" json:"share,omitempty"`
}

type Share struct {
	Entry       []Child   `xml:"entry" json:"entry,omitempty"`
	ID          string    `xml:"id,attr" json:"id"`
	URL         string    `xml:"url,attr" json:"url"`
	Description *string   `xml:"description,attr,omitempty" json:"description,omitempty"`
	Username    string    `xml:"username,attr" json:"username"`
	Created     DateTime  `xml:"created,attr" json:"created"`
	Expires     *DateTime `xml:"expires,attr,omitempty" json:"expires,omitempty"`
	LastVisited *DateTime `xml:"lastVisited,attr,omitempty" json:"lastVisited,omitempty"`
	VisitCount  int       `xml:"visitCount,attr" json:"visitCount"`
}

type Starred struct {
	Artist []Artist `xml:"artist" json:"artist,omitempty"`
	Album  []Child  `xml:"album" json:"album,omitempty"`
	Song   []Child  `xml:"song" json:"song,omitempty"`
}

type AlbumInfo struct {
	Notes          *string `xml:"notes,omitempty" json:"notes,omitempty"`
	MusicBrainzID  *string `xml:"musicBrainzId,omitempty" json:"musicBrainzId,omitempty"`
	LastFmURL      *string `xml:"lastFmUrl,omitempty" json:"lastFmUrl,omitempty"`
	SmallImageURL  *string `xml:"smallImageUrl,omitempty" json:"smallImageUrl,omitempty"`
	MediumImageURL *string `xml:"mediumImageUrl,omitempty" json:"mediumImageUrl,omitempty"`
	LargeImageURL  *string `xml:"largeImageUrl,omitempty" json:"largeImageUrl,omitempty"`
}

type ArtistInfoBase struct {
	Biography      *string `xml:"biography,omitempty" json:"biography,omitempty"`
	MusicBrainzID  *string `xml:"musicBrainzId,omitempty" json:"musicBrainzId,omitempty"`
	LastFmURL      *string `xml:"lastFmUrl,omitempty" json:"lastFmUrl,omitempty"`
	SmallImageURL  *string `xml:"smallImageUrl,omitempty" json:"smallImageUrl,omitempty"`
	MediumImageURL *string `xml:"mediumImageUrl,omitempty" json:"mediumImageUrl,omitempty"`
	LargeImageURL  *string `xml:"largeImageUrl,omitempty" json:"largeImageUrl,omitempty"`
}

type ArtistInfo struct {
	ArtistInfoBase
	SimilarArtist []Artist `xml:"similarArtist" json:"similarArtist,omitempty"`
}

type ArtistInfo2 struct {
	ArtistInfoBase
	SimilarArtist []ArtistID3 `xml:"similarArtist" json:"similarArtist,omitempty"`
}

type SimilarSongs struct {
	Song []Child `xml:"song" json:"song,omitempty"`
}

type SimilarSongs2 struct {
	Song []Child `xml:"song" json:"song,omitempty"`
}

type TopSongs struct {
	Song []Child `xml:"song" json:"song,omitempty"`
}

type Starred2 struct {
	Artist []ArtistID3 `xml:"artist" json:"artist,omitempty"`
	Album  []AlbumID3  `xml:"album" json:"album,omitempty"`
	Song   []Child     `xml:"song" json:"song,omitempty"`
}

type License struct {
	Valid          bool      `xml:"valid,attr" json:"valid"`
	Email          *string   `xml:"email,attr,omitempty" json:"email,omitempty"`
	LicenseExpires *DateTime `xml:"licenseExpires,attr,omitempty" json:"licenseExpires,omitempty"`
	TrialExpires   *DateTime `xml:"trialExpires,attr,omitempty" json:"trialExpires,omitempty"`
}

type ScanStatus struct {
	Scanning bool   `xml:"scanning,attr" json:"scanning"`
	Count    *int64 `xml:"count,attr,omitempty" json:"count,omitempty"`
}

type Users struct {
	User []User `xml:"user" json:"user,omitempty"`
}

type User struct {
	Folder              []int     `xml:"folder" json:"folder,omitempty"`
	Username            string    `xml:"username,attr" json:"username"`
	Email               *string   `xml:"email,attr,omitempty" json:"email,omitempty"`
	ScrobblingEnabled   bool      `xml:"scrobblingEnabled,attr" json:"scrobblingEnabled"`
	MaxBitRate          *int      `xml:"maxBitRate,attr,omitempty" json:"maxBitRate,omitempty"`
	AdminRole           bool      `xml:"adminRole,attr" json:"adminRole"`
	SettingsRole        bool      `xml:"settingsRole,attr" json:"settingsRole"`
	DownloadRole        bool      `xml:"downloadRole,attr" json:"downloadRole"`
	UploadRole          bool      `xml:"uploadRole,attr" json:"uploadRole"`
	PlaylistRole        bool      `xml:"playlistRole,attr" json:"playlistRole"`
	CoverArtRole        bool      `xml:"coverArtRole,attr" json:"coverArtRole"`
	CommentRole         bool      `xml:"commentRole,attr" json:"commentRole"`
	PodcastRole         bool      `xml:"podcastRole,attr" json:"podcastRole"`
	StreamRole          bool      `xml:"streamRole,attr" json:"streamRole"`
	JukeboxRole         bool      `xml:"jukeboxRole,attr" json:"jukeboxRole"`
	ShareRole           bool      `xml:"shareRole,attr" json:"shareRole"`
	VideoConversionRole bool      `xml:"videoConversionRole,attr" json:"videoConversionRole"`
	AvatarLastChanged   *DateTime `xml:"avatarLastChanged,attr,omitempty" json:"avatarLastChanged,omitempty"`
}

type Error struct {
	Code    int     `xml:"code,attr" json:"code"`
	Message *string `xml:"message,attr,omitempty" json:"message,omitempty"`
}

package model

type Album struct {
	AlbumID           int64       `json:"albumId"`         // 专辑id
	AlbumName         string      `json:"albumName"`       // 专辑名字
	Artist            interface{} `json:"artist"`          // 歌手
	ArtistAvatarURL   string      `json:"artistAvatarUrl"` // 歌手头像
	ArtistID          int64       `json:"artistId"`        // 歌手id
	ArtistInfoList    interface{} `json:"artistInfoList"`
	ArtistName        string      `json:"artistName"` // 歌手名字
	ArtistNames       string      `json:"artistNames"`
	BlurImgURL        string      `json:"blurImgUrl"`
	Bundling          int64       `json:"bundling"`
	CoverURL          string      `json:"coverUrl"` // 封面url
	CustomPriceConfig interface{} `json:"customPriceConfig"`
	LiveType          int64       `json:"liveType"`
	RoomNo            int64       `json:"roomNo"`
	Stars             bool        `json:"stars"`
}

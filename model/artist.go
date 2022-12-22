package model

type Artist struct {
	AlbumSize   int64    `json:"albumSize"` // 专辑数
	BriefDesc   string   `json:"briefDesc"` // 描述
	Cover       string   `json:"cover"`     // 主页图
	ID          int64    `json:"id"`        // 歌手id
	IdentifyTag []string `json:"identifyTag"`
	Identities  []string `json:"identities"`
	MusicSize   int64    `json:"musicSize"`  // 音乐数量
	MvSize      int64    `json:"mvSize"`     // mv数量
	Name        string   `json:"name"`       // 歌手名字
	TransNames  []string `json:"transNames"` // 翻译名
}

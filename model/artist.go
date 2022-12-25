package model

type Artist struct {
	AccountID   int64    `json:"accountId"`
	AlbumSize   int64    `json:"albumSize"` // 专辑数
	Alias       []string `json:"alias"`
	BriefDesc   string   `json:"briefDesc"` // 描述
	Cover       string   `json:"cover"`     // 主页图
	Followed    bool     `json:"followed"`
	ID          int64    `json:"id"` // 歌手id
	IdentifyTag []string `json:"identifyTag"`
	Identities  []string `json:"identities"`
	Img1V1ID    int64    `json:"img1v1Id"`
	Img1V1URL   string   `json:"img1v1Url"`
	MusicSize   int64    `json:"musicSize"` // 音乐数量
	MvSize      int64    `json:"mvSize"`    // mv数量
	Name        string   `json:"name"`      // 歌手名字
	PicID       int64    `json:"picId"`
	PicURL      string   `json:"picUrl"`
	PublishTime int64    `json:"publishTime"`
	Rank        Rank     `json:"rank"`
	TopicPerson int64    `json:"topicPerson"`
	Trans       string   `json:"trans"`
	TransNames  []string `json:"transNames"` // 翻译名
}

// rank.go

type Rank struct {
	Rank int64 `json:"rank"`
	Type int64 `json:"type"`
}

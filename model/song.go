package model

type Song struct {
	Al              Album    `json:"al"`   // 专辑
	Alia            []string `json:"alia"` // 别名
	Ar              []Artist `json:"ar"`   // 歌手列表
	CD              string   `json:"cd"`
	Copyright       int64    `json:"copyright"`
	DjID            int64    `json:"djId"` // 是否dj节目
	Dt              int64    `json:"dt"`   // 歌曲时长
	Fee             int64    `json:"fee"`  // 歌曲播放权限
	ID              int64    `json:"id"`   // 歌曲id
	Mark            int64    `json:"mark"`
	Mv              int64    `json:"mv"`   // mvid
	Name            string   `json:"name"` // 歌曲名
	No              int64    `json:"no"`
	OriginCoverType int64    `json:"originCoverType"` // 翻唱类型
	Pop             int64    `json:"pop"`             // 歌曲热度
	PublishTime     int64    `json:"publishTime"`     // 发行时间
	RtUrls          []string `json:"rtUrls"`
	SID             int64    `json:"s_id"`
	Single          int64    `json:"single"` // 有无专辑
	T               int64    `json:"t"`      // 歌曲来源
}

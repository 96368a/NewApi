package model

type Playlist struct {
	AdType             int64       `json:"adType"`
	BackgroundCoverID  int64       `json:"backgroundCoverId"`
	BackgroundCoverURL interface{} `json:"backgroundCoverUrl"`
	CloudTrackCount    int64       `json:"cloudTrackCount"`
	CommentCount       int64       `json:"commentCount"`
	CommentThreadID    string      `json:"commentThreadId"`
	Copied             bool        `json:"copied"`
	CoverImgID         int64       `json:"coverImgId"`
	CoverImgIDStr      string      `json:"coverImgId_str"`
	CoverImgURL        string      `json:"coverImgUrl"` // 封面图片url
	CreateTime         int64       `json:"createTime"`  // 创建时间戳
	Creator            User        `json:"creator"`
	Description        string      `json:"description"`
	GradeStatus        string      `json:"gradeStatus"`
	HighQuality        bool        `json:"highQuality"`
	ID                 int64       `json:"id"`   // id
	Name               string      `json:"name"` // 播放列表名字
	NewImported        bool        `json:"newImported"`
	OpRecommend        bool        `json:"opRecommend"`
	Ordered            bool        `json:"ordered"`
	PlayCount          int64       `json:"playCount"` // 播放数
	Privacy            int64       `json:"privacy"`
	ShareCount         int64       `json:"shareCount"`
	SpecialType        int64       `json:"specialType"`
	Status             int64       `json:"status"`
	Subscribed         bool        `json:"subscribed"`
	SubscribedCount    int64       `json:"subscribedCount"`
	//Subscribers           []Subscriber     `json:"subscribers"`
	Tags                  []string         `json:"tags"`
	TitleImage            int64            `json:"titleImage"`
	TrackCount            int64            `json:"trackCount"`
	TrackIDS              []TrackIDElement `json:"trackIds"` // 歌单歌曲id
	TrackNumberUpdateTime int64            `json:"trackNumberUpdateTime"`
	//Tracks                []Song           `json:"tracks"` // 歌单歌曲内容
	TrackUpdateTime int64       `json:"trackUpdateTime"`
	UpdateTime      int64       `json:"updateTime"`
	UserID          int64       `json:"userId"` // 用户id
	VideoIDS        interface{} `json:"videoIds"`
	Videos          interface{} `json:"videos"`
}

type TrackIDElement struct {
	At  int64 `json:"at"`
	ID  int64 `json:"id"`
	Uid int64 `json:"uid"`
	V   int64 `json:"v"`
}

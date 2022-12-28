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
	Tracks                []Song           `json:"tracks"` // 歌单歌曲内容
	TrackUpdateTime       int64            `json:"trackUpdateTime"`
	UpdateTime            int64            `json:"updateTime"`
	UserID                int64            `json:"userId"` // 用户id
	VideoIDS              interface{}      `json:"videoIds"`
	Videos                interface{}      `json:"videos"`
}

type TrackIDElement struct {
	Alg        interface{} `json:"alg"`
	At         int64       `json:"at"`
	F          interface{} `json:"f"`
	ID         int64       `json:"id"`
	RcmdReason string      `json:"rcmdReason"`
	Sc         interface{} `json:"sc"`
	Sr         interface{} `json:"sr"`
	T          int64       `json:"t"`
	Uid        int64       `json:"uid"`
	V          int64       `json:"v"`
}

// creator.go

// subscriber.go

type Subscriber struct {
	AccountStatus            int64       `json:"accountStatus"`
	Anchor                   bool        `json:"anchor"`
	AuthenticationTypes      int64       `json:"authenticationTypes"`
	Authority                int64       `json:"authority"`
	AuthStatus               int64       `json:"authStatus"`
	AvatarDetail             interface{} `json:"avatarDetail"`
	AvatarImgID              int64       `json:"avatarImgId"`
	SubscriberAvatarImgIDStr string      `json:"avatarImgId_str"`
	AvatarImgIDStr           string      `json:"avatarImgIdStr"`
	AvatarURL                string      `json:"avatarUrl"`
	BackgroundImgID          int64       `json:"backgroundImgId"`
	BackgroundImgIDStr       string      `json:"backgroundImgIdStr"`
	BackgroundURL            string      `json:"backgroundUrl"`
	Birthday                 int64       `json:"birthday"`
	City                     int64       `json:"city"`
	DefaultAvatar            bool        `json:"defaultAvatar"`
	Description              string      `json:"description"`
	DetailDescription        string      `json:"detailDescription"`
	DjStatus                 int64       `json:"djStatus"`
	Experts                  interface{} `json:"experts"`
	ExpertTags               interface{} `json:"expertTags"`
	Followed                 bool        `json:"followed"`
	Gender                   int64       `json:"gender"`
	Mutual                   bool        `json:"mutual"`
	Nickname                 string      `json:"nickname"`
	Province                 int64       `json:"province"`
	RemarkName               interface{} `json:"remarkName"`
	Signature                string      `json:"signature"`
	UserID                   int64       `json:"userId"`
	UserType                 int64       `json:"userType"`
	VipType                  int64       `json:"vipType"`
}

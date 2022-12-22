package model

type Album struct {
	Alias           []string    `json:"alias"`
	Artist          Artist      `json:"artist"`
	Artists         []Artist    `json:"artists"`
	BlurPicURL      string      `json:"blurPicUrl"`
	BriefDesc       string      `json:"briefDesc"`
	CommentThreadID string      `json:"commentThreadId"`
	Company         interface{} `json:"company"`
	CompanyID       int64       `json:"companyId"`
	CopyrightID     int64       `json:"copyrightId"`
	Description     string      `json:"description"`
	ID              int64       `json:"id"`
	Info            Info        `json:"info"`
	Mark            int64       `json:"mark"`
	Name            string      `json:"name"`
	OnSale          bool        `json:"onSale"`
	Paid            bool        `json:"paid"`
	Pic             int64       `json:"pic"`
	PicID           int64       `json:"picId"`
	PicURL          string      `json:"picUrl"`
	PublishTime     int64       `json:"publishTime"`
	Size            int64       `json:"size"`
	Songs           []string    `json:"songs"`
	Status          int64       `json:"status"`
	SubType         string      `json:"subType"`
	Tags            string      `json:"tags"`
	Type            string      `json:"type"`
}

type Info struct {
	CommentCount     int64         `json:"commentCount"`
	Comments         interface{}   `json:"comments"`
	CommentThread    CommentThread `json:"commentThread"`
	LatestLikedUsers interface{}   `json:"latestLikedUsers"`
	Liked            bool          `json:"liked"`
	LikedCount       int64         `json:"likedCount"`
	ResourceID       int64         `json:"resourceId"`
	ResourceType     int64         `json:"resourceType"`
	ShareCount       int64         `json:"shareCount"`
	ThreadID         string        `json:"threadId"`
}

type CommentThread struct {
	CommentCount     int64        `json:"commentCount"`
	HotCount         int64        `json:"hotCount"`
	ID               string       `json:"id"`
	LatestLikedUsers interface{}  `json:"latestLikedUsers"`
	LikedCount       int64        `json:"likedCount"`
	ResourceID       int64        `json:"resourceId"`
	ResourceInfo     ResourceInfo `json:"resourceInfo"`
	ResourceOwnerID  int64        `json:"resourceOwnerId"`
	ResourceTitle    string       `json:"resourceTitle"`
	ResourceType     int64        `json:"resourceType"`
	ShareCount       int64        `json:"shareCount"`
}

type ResourceInfo struct {
	Creator   interface{} `json:"creator"`
	EncodedID interface{} `json:"encodedId"`
	ID        int64       `json:"id"`
	ImgURL    string      `json:"imgUrl"`
	Name      string      `json:"name"`
	SubTitle  interface{} `json:"subTitle"`
	UserID    int64       `json:"userId"`
	WebURL    interface{} `json:"webUrl"`
}

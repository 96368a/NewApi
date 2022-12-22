package model

//type User struct {
//	ID       uint64 `json:"id"`       // 用户id
//	Username string `json:"username"` // 用户名
//	Nickname string `json:"nickname"` // 昵称
//
//	Signature string `json:"signature"` // 个性签名
//	IsAdmin   int    `json:"isAdmin"`
//}

type User struct {
	AccountStatus         int64       `json:"accountStatus"`
	Anchor                bool        `json:"anchor"`
	AuthenticationTypes   int64       `json:"authenticationTypes"`
	Authority             int64       `json:"authority"`
	AuthStatus            int64       `json:"authStatus"`
	AvatarDetail          interface{} `json:"avatarDetail"`
	AvatarImgID           int64       `json:"avatarImgId"`
	CreatorAvatarImgIDStr string      `json:"avatarImgId_str"`
	AvatarImgIDStr        string      `json:"avatarImgIdStr"`
	AvatarURL             string      `json:"avatarUrl"`
	BackgroundImgID       int64       `json:"backgroundImgId"`
	BackgroundImgIDStr    string      `json:"backgroundImgIdStr"`
	BackgroundURL         string      `json:"backgroundUrl"`
	Birthday              int64       `json:"birthday"`
	City                  int64       `json:"city"`
	DefaultAvatar         bool        `json:"defaultAvatar"`
	Description           string      `json:"description"`
	DetailDescription     string      `json:"detailDescription"`
	DjStatus              int64       `json:"djStatus"`
	Experts               interface{} `json:"experts"`
	ExpertTags            interface{} `json:"expertTags"`
	Followed              bool        `json:"followed"`
	Gender                int64       `json:"gender"`
	Mutual                bool        `json:"mutual"`
	Nickname              string      `json:"nickname"` //用户昵称
	Password              string      `json:"password"` // 用户密码
	Phone                 string      `json:"phone"`    //用户手机好
	Email                 string      `json:"email"`    //用户邮箱
	Province              int64       `json:"province"`
	RemarkName            interface{} `json:"remarkName"`
	Signature             string      `json:"signature"`
	UserID                int64       `json:"userId"`
	UserType              int64       `json:"userType"`
	VipType               int64       `json:"vipType"`
	Status                int         `json:"status"`
}

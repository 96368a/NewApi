package model

type User struct {
	ID        uint64 `json:"id"`        // 用户id
	Username  string `json:"username"`  // 用户名
	Nickname  string `json:"nickname"`  // 昵称
	Password  string `json:"password"`  // 用户密码
	Signature string `json:"signature"` // 个性签名
	IsAdmin   int    `json:"isAdmin"`
}

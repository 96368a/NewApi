package services

import (
	"encoding/json"
	"fmt"
	"github.com/96368a/NewApi/model"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAdd(t *testing.T) {
	user := model.User{
		AccountStatus:         0,
		Anchor:                false,
		AuthenticationTypes:   0,
		Authority:             0,
		AuthStatus:            0,
		AvatarDetail:          nil,
		AvatarImgID:           0,
		CreatorAvatarImgIDStr: "",
		AvatarImgIDStr:        "",
		AvatarURL:             "",
		BackgroundImgID:       0,
		BackgroundImgIDStr:    "",
		BackgroundURL:         "",
		Birthday:              0,
		City:                  0,
		DefaultAvatar:         false,
		Description:           "",
		DetailDescription:     "",
		DjStatus:              0,
		Experts:               nil,
		ExpertTags:            nil,
		Followed:              false,
		Gender:                0,
		Mutual:                false,
		Nickname:              "测试用户1",
		Password:              "123456",
		Phone:                 "18888888888",
		Email:                 "test2@qq.com",
		Province:              0,
		RemarkName:            nil,
		Signature:             "",
		UserID:                10001,
		UserType:              0,
		VipType:               0,
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)
	i := Add(user)
	fmt.Println(i)
}

func TestGetByPhone(t *testing.T) {
	phone := GetOneByPhone("15577891000")
	fmt.Println(phone)
	fmt.Println("1")
}

type Message struct {
	code int  `json:"code"`
	more bool `json:"more"`
	size int  `json:"size"`
	//followeds []model.User `json:"followeds"`
}

func TestAddUsers(t *testing.T) {
	resp, err := http.Get("https://y.233c.cn/user/followeds?uid=32953014")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	//message := Message{}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("反序列化失败", err)
	}

	var user model.User

	datas := data["followeds"].([]interface{})
	for _, v := range datas {
		data2, _ := json.Marshal(v)
		json.Unmarshal(data2, &user)
		fmt.Printf("%+v\n", user)
		user.Password = "123456"
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashPassword)
		user.Phone = fmt.Sprintf("%d", 18888888888-user.UserID)
		user.Email = fmt.Sprintf("%d%s", user.UserID, "@qq.com")
		Add(user)
		fmt.Printf("%v", user)
	}
}

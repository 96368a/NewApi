package services

import (
	"fmt"
	"github.com/96368a/NewApi/model"
	"golang.org/x/crypto/bcrypt"
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

package controller

import (
	"github.com/96368a/NewApi/dto"
	"github.com/96368a/NewApi/model"
	"github.com/96368a/NewApi/services"
	"github.com/96368a/NewApi/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

func LoginByPhone(c *gin.Context) {
	var user dto.UserDto
	c.ShouldBind(&user)
	phone := c.Query("phone")
	password := c.Query("password")
	if phone != "" {
		user.Phone = phone
	}
	if password != "" {
		user.Password = password
	}
	if len(user.Phone) < 4 || len(user.Password) < 4 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	user1 := services.GetOneByPhone(user.Phone)
	if user1 == nil || user1.UserID == 0 {
		utils.Fail(c, http.StatusBadRequest, "用户不存在", nil)
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user1.Password), []byte(user.Password))
	if err != nil {
		utils.Fail(c, http.StatusUnauthorized, "密码错误", nil)
		return
	}
	token, err := utils.ReleaseToken(*user1)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "token生成失败", nil)
	}
	c.Header("Authorization", token)
	user1.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"account": user1,
		"token":   token,
		"profile": user1,
	})
}

func LoginByEmail(c *gin.Context) {
	var user dto.UserDto
	c.ShouldBind(&user)
	email := c.Query("email")
	password := c.Query("password")
	if email != "" {
		user.Email = email
	}
	if password != "" {
		user.Password = password
	}

	if len(user.Email) < 4 || len(user.Password) < 4 {
		utils.Fail(c, http.StatusBadRequest, "参数错误", nil)
		return
	}
	user1 := services.GetOneByEmail(user.Email)
	if user1 == nil || user1.UserID == 0 {
		utils.Fail(c, http.StatusBadRequest, "用户不存在", nil)
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user1.Password), []byte(user.Password))
	if err != nil {
		utils.Fail(c, http.StatusUnauthorized, "密码错误", nil)
		return
	}
	token, err := utils.ReleaseToken(*user1)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "token生成失败", nil)
	}
	c.Header("Authorization", token)
	//c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	user1.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"account": user1,
		"token":   token,
		"profile": user1,
	})
}

func LoginStatus(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "token格式错误",
		})
		return
	}

	tokenString = tokenString[7:]
	token, claims, err := utils.ParseToken(tokenString)

	var userId int64

	if err != nil || !token.Valid {
		userId = -1
	} else {
		userId = claims.UserId
	}
	//获取用户信息
	var user *model.User
	user = services.GetOne(userId)
	//fmt.Printf("%v id: %v\v", user, userId)
	// 验证用户是否存在
	if user.UserID == 0 {
		user.Status = -10
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"code":    200,
				"account": user,
				"profile": nil,
			},
		})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"code":    200,
			"account": user,
			"profile": user,
		},
	})
}

//func UserInfo(c *gin.Context) {
//	user, _ := c.Get("user")
//	utils.Success(c, gin.H{
//		"user": vo.ToUserVO(user.(model.User)),
//	}, "获取成功")
//}

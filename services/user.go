package services

import (
	"context"
	"github.com/96368a/NewApi/model"
	"go.mongodb.org/mongo-driver/bson"
)

func AddUser(user model.User) interface{} {
	one := GetOneUser(user.UserID)
	if one.UserID == user.UserID {
		return nil
	}
	one = GetOneUserByEmail(user.Email)
	if one.Email == user.Email {
		return nil
	}
	one = GetOneUserByPhone(user.Phone)
	if one.Phone == user.Phone {
		return nil
	}
	result, err := model.UserCol.InsertOne(context.TODO(), user)
	if err != nil {
		return nil
	}
	return result.InsertedID
}

func GetOneUser(id int64) *model.User {
	var user model.User
	model.UserCol.FindOne(context.TODO(), bson.D{{"userid", id}}).Decode(&user)
	return &user
}

func GetOneUserByEmail(email string) *model.User {
	var user model.User
	model.UserCol.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	return &user
}

func GetOneUserByPhone(phone string) *model.User {
	var user model.User
	model.UserCol.FindOne(context.TODO(), bson.D{{"phone", phone}}).Decode(&user)
	return &user
}

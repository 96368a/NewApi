package services

import (
	"context"
	"github.com/96368a/NewApi/model"
	"go.mongodb.org/mongo-driver/bson"
)

func Add(user model.User) interface{} {
	one := GetOne(user.UserID)
	if one.UserID == user.UserID {
		return nil
	}
	one = GetOneByEmail(user.Email)
	if one.Email == user.Email {
		return nil
	}
	one = GetOneByPhone(user.Phone)
	if one.Phone == user.Phone {
		return nil
	}
	result, err := model.UserCol.InsertOne(context.TODO(), user)
	if err != nil {
		return nil
	}
	return result.InsertedID
}

func GetOne(id int64) *model.User {
	var user model.User
	model.UserCol.FindOne(context.TODO(), bson.D{{"userid", id}}).Decode(&user)
	return &user
}

func GetOneByEmail(email string) *model.User {
	var user model.User
	model.UserCol.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	return &user
}

func GetOneByPhone(phone string) *model.User {
	var user model.User
	model.UserCol.FindOne(context.TODO(), bson.D{{"phone", phone}}).Decode(&user)
	return &user
}
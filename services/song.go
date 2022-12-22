package services

import (
	"context"
	"github.com/96368a/NewApi/model"
	"go.mongodb.org/mongo-driver/bson"
)

func AddSong(song model.Song) interface{} {
	one, err := model.SongCol.InsertOne(context.TODO(), song)
	if err != nil {
		return nil
	}
	return one.InsertedID

}
func GetOneSongByObjectId(id int64) *model.Song {
	var song model.Song
	model.SongCol.FindOne(context.TODO(), bson.D{{"songid", id}}).Decode(&song)
	return &song
}

func GetOneSong(id int64) *model.Song {
	var song model.Song
	model.SongCol.FindOne(context.TODO(), bson.D{{"songid", id}}).Decode(&song)
	return &song
}

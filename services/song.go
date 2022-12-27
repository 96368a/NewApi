package services

import (
	"context"
	"github.com/96368a/NewApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddSong(song model.Song) interface{} {
	one := GetOneSong(song.ID)
	if one != nil {
		return nil
	}
	one1, err := model.SongCol.InsertOne(context.TODO(), song)
	if err != nil {
		return nil
	}
	return one1.InsertedID

}
func GetOneSongByObjectId(id int64) *model.Song {
	var song model.Song
	model.SongCol.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&song)
	if song.ID == 0 {
		return nil
	}
	return &song
}

func GetOneSong(id int64) *model.Song {
	var song model.Song
	model.SongCol.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&song)
	if song.ID == 0 {
		return nil
	}
	return &song
}

func SearchSong(keyword string, limit int64, offset int64) []*model.Song {
	var songs []*model.Song
	cursor, err := model.SongCol.Find(context.TODO(), bson.D{{"name", bson.D{{"$regex", keyword}}}}, options.Find().SetSkip(offset).SetLimit(limit))
	if err != nil {
		return nil
	}
	if cursor.All(context.TODO(), &songs) != nil {
		return nil
	}
	if len(songs) == 0 {
		return nil
	}
	return songs
}

func GetSongs(ids []int64) []*model.Song {
	var songs []*model.Song
	cursor, err := model.SongCol.Find(context.TODO(), bson.D{{"id", bson.D{{"$in", ids}}}})
	if err != nil {
		return nil
	}
	if cursor.All(context.TODO(), &songs) != nil {
		return nil
	}
	if len(songs) == 0 {
		return nil
	}
	return songs
}

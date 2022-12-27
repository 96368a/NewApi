package services

import (
	"context"
	"github.com/96368a/NewApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddPlaylist(playlist model.Playlist) interface{} {
	one := GetOnePlaylist(playlist.ID)
	if one != nil {
		return nil
	}
	one1, err := model.PlaylistCol.InsertOne(context.TODO(), playlist)
	if err != nil {
		return nil
	}
	return one1.InsertedID
}

func GetOnePlaylist(id int64) *model.Playlist {
	var playlist model.Playlist
	model.PlaylistCol.FindOne(context.TODO(), model.Playlist{ID: id}).Decode(&playlist)
	if playlist.ID == 0 {
		return nil
	}
	return &playlist
}

func GetOnePlaylistByObjectId(id int64) *model.Playlist {
	var playlist model.Playlist
	model.PlaylistCol.FindOne(context.TODO(), model.Playlist{ID: id}).Decode(&playlist)
	if playlist.ID == 0 {
		return nil
	}
	return &playlist
}

func SearchPlaylist(keyword string, limit int64, offset int64) []*model.Playlist {
	var playlists []*model.Playlist
	cursor, err := model.PlaylistCol.Find(context.TODO(), bson.D{{"name", bson.D{{"$regex", keyword}}}}, options.Find().SetSkip(offset).SetLimit(limit))
	if err != nil {
		return nil
	}
	if cursor.All(context.TODO(), &playlists) != nil {
		return nil
	}
	if len(playlists) == 0 {
		return nil
	}
	return playlists
}

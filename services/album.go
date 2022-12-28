package services

import (
	"context"
	"github.com/96368a/NewApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddAlbum(album model.Album) interface{} {
	one := GetOneAlbum(album.ID)
	if one != nil {
		return nil
	}
	one1, err := model.AlbumCol.InsertOne(context.TODO(), album)
	if err != nil {
		return nil
	}
	return one1.InsertedID
}

func GetOneAlbum(id int64) *model.Album {
	var album model.Album
	model.AlbumCol.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&album)
	if album.ID == 0 {
		return nil
	}
	return &album
}

func GetOneAlbumByObjectId(id int64) *model.Album {
	var album model.Album
	model.AlbumCol.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&album)
	if album.ID == 0 {
		return nil
	}
	return &album
}

func SearchAlbum(keyword string, limit int64, offset int64) []*model.Album {
	var albums []*model.Album
	cursor, err := model.AlbumCol.Find(context.TODO(), bson.D{{"name", bson.D{{"$regex", keyword}}}}, options.Find().SetSkip(offset).SetLimit(limit))
	if err != nil {
		return nil
	}
	if cursor.All(context.TODO(), &albums) != nil {
		return nil
	}
	if len(albums) == 0 {
		return nil
	}
	return albums
}

func GetAlbumSongs(id int64, limit int) []*model.Song {
	var songs []*model.Song
	cursor, err := model.SongCol.Find(context.TODO(), bson.D{{"al.id", id}}, options.Find().SetLimit(int64(limit)))
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

func GetNewAlbum(limit int64, offset int64) []*model.Album {
	var albums []*model.Album
	cursor, err := model.AlbumCol.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{"publishTime", -1}}).SetLimit(limit).SetSkip(offset))
	if err != nil {
		return nil
	}
	if cursor.All(context.TODO(), &albums) != nil {
		return nil
	}
	if len(albums) == 0 {
		return nil
	}
	return albums
}

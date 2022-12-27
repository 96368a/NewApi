package services

import (
	"context"
	"github.com/96368a/NewApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddArtist(artist model.Artist) interface{} {
	one := GetOneArtist(artist.ID)
	if one != nil {
		return nil
	}
	one1, err := model.ArtistCol.InsertOne(context.TODO(), artist)
	if err != nil {
		return nil
	}
	return one1.InsertedID
}

func GetOneArtist(id int64) *model.Artist {
	var artist model.Artist
	model.ArtistCol.FindOne(context.TODO(), model.Artist{ID: id}).Decode(&artist)
	if artist.ID == 0 {
		return nil
	}
	return &artist
}

func GetOneArtistByObjectId(id int64) *model.Artist {
	var artist model.Artist
	model.ArtistCol.FindOne(context.TODO(), model.Artist{ID: id}).Decode(&artist)
	if artist.ID == 0 {
		return nil
	}
	return &artist
}

func SearchArtist(keyword string, limit int64, offset int64) []*model.Artist {
	var artists []*model.Artist
	cursor, err := model.ArtistCol.Find(context.TODO(), bson.D{{"name", bson.D{{"$regex", keyword}}}}, options.Find().SetSkip(offset).SetLimit(limit))
	if err != nil {
		return nil
	}
	if cursor.All(context.TODO(), &artists) != nil {
		return nil
	}
	if len(artists) == 0 {
		return nil
	}
	return artists
}

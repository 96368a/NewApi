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
	model.ArtistCol.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&artist)
	if artist.ID == 0 {
		return nil
	}
	return &artist
}

func GetOneArtistByObjectId(id int64) *model.Artist {
	var artist model.Artist
	model.ArtistCol.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&artist)
	if artist.ID == 0 {
		return nil
	}
	return &artist
}

func GetArtistAlbum(id int64, limit int) []*model.Album {
	var albums []*model.Album
	cursor, err := model.AlbumCol.Find(context.TODO(), bson.D{{"artist.id", id}}, options.Find().SetLimit(int64(limit)))
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

//func GetArtist(id int64, limit int) *model.Artist {
//
//}

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

func GetArtistHotSongs(id int64) []*model.Song {
	var songs []*model.Song
	cursor, err := model.SongCol.Find(context.TODO(), bson.D{{"ar.id", id}}, options.Find().SetLimit(10))
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

func GetArtistSongs(id int64, limit int64, offset int64) []*model.Song {
	var songs []*model.Song
	cursor, err := model.SongCol.Find(context.TODO(), bson.D{{"ar.id", id}}, options.Find().SetLimit(limit).SetSkip(offset))
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

func GetArtistToplist(type1 int) []*model.Artist {
	var artists []*model.Artist
	cursor, err := model.ArtistCol.Find(context.TODO(), bson.D{{"rank.type", type1}}, options.Find().SetSort(bson.D{{"rank.rank", 1}}).SetLimit(100))
	if err != nil {
		return nil
	}
	if cursor.All(context.TODO(), &artists) != nil {
		return nil
	}
	return artists

}

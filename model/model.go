package model

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCol *mongo.Collection
var SongCol *mongo.Collection
var AlbumCol *mongo.Collection
var ArtistCol *mongo.Collection

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	//defer func() {
	//	if err := client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()
	UserCol = client.Database("music").Collection("user")
	SongCol = client.Database("music").Collection("song")
	AlbumCol = client.Database("music").Collection("album")
	ArtistCol = client.Database("music").Collection("artist")

}

package model

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MusicCol *mongo.Collection
var UserCol *mongo.Collection

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
	MusicCol = client.Database("music").Collection("music")
	UserCol = client.Database("music").Collection("user")
	//title := "Back to the Future"
	//var result bson.M
	//err = MusicCol.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	//if err == mongo.ErrNoDocuments {
	//	fmt.Printf("No document was found with the title %s\n", title)
	//	return
	//}
	//if err != nil {
	//	panic(err)
	//}
	//jsonData, err := json.MarshalIndent(result, "", "    ")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", jsonData)
}

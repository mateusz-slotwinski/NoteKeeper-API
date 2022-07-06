package database

import (
	context "context"
	fmt "fmt"
	time "time"

	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"

	Config "NoteKeeperAPI/src/config"
	Helpers "NoteKeeperAPI/src/helpers"
)

const (
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

func ConnectDB() *mongo.Client {
	Config.Init()

	username := Config.DB_Username
	password := Config.DB_Password
	clusterEndpoint := Config.DB_ClusterEndpoint

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	Helpers.PrintError(err)

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	Helpers.PrintError(err)

	err = client.Ping(ctx, nil)
	Helpers.PrintError(err)

	if Config.Mode == "DEV" {
		fmt.Println("Connected to MongoDB")
	}
	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(Config.DB_Collection).Collection(collectionName)
	return collection
}

// var Client, _ = mongo.NewClient(options.Client().ApplyURI(Config.DBAdress))
// var Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
// var db = Client.Database("mo1394_lightning")
// var links = db.Collection("links")

// func Init() {
// 	var err = Client.Connect(Ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)
var Client *mongo.Client
func ConnectDB(uri string){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

	client,err := mongo.Connect(options.Client().ApplyURI(uri))
	if err!=nil{
      log.Fatal(err)
	}
	   if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("Could not connect to MongoDB:", err)
    }

	log.Println("Connected to MongoDB")
	Client = client

} 
func GetCollection(dbName,collName string) *mongo.Collection{
	if Client == nil {
		log.Fatal("MongoDB not connected")
	}
	return Client.Database(dbName).Collection(collName)
}
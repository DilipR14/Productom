package database

import(
	"context"
	"log"
	"fmt"
	"time"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DataBaseSet() *mongo.Client{
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err! =nil {
		log. Fatal(err)
	}

	ctx, cancel := context.withTimeOut(context.Background(), 10*time.second)

	defer cancel()

	err = client.Connect(ctx)

	if err! = nil{
		log.Fatal(err)
	}

	err = client.Ping(context. TODO(),nil)
	if err!= nil{
		log.Println("failed to connect to mongodb")
		return nil
	}

	fmt.Println("Successfully connected to mongodb")
	return client

}

func UserData(client *mongo.client, collectionName string) *mongo.collection{
	var collection *mongo.collection = client.DataBase("Ecommerce").collection(collectionName)
	return collection
}

func ProductData(client *mongo.client, collectionName string) *mongo.collection{
	var productData *mongo.collection = client.DataBase("Ecommerce").collection(collectionName)
	return productcollection
}
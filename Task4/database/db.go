package database

import (
	"context"
	"fmt"
	"task4/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//const connectionString string = "mongodb://localhost:27017"

const connectionString string = "mongodb://mongo1:27017,mongo2:27017,mongo3:27017/?replicaSet=dbrs"

var Client *mongo.Client

func ConnectDB() error {
	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Println("Database Error: ", err.Error())
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = Client.Connect(ctx)
	if err != nil {
		fmt.Println("Database Connection Error: ", err.Error())
		return err
	}
	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return err
	}
	return nil
}

func InsertData(data models.Data) error {
	err := ConnectDB()
	if err != nil {
		return err
	}
	JsonDataCollection := Client.Database("testMongo").Collection("jsonData")
	// data.Title = "test"
	filter := bson.D{{"id", data.Id}}
	updates := bson.D{{"$set", data}}
	upsertFlag := options.Update().SetUpsert(true)
	jsonDataResult, err := JsonDataCollection.UpdateOne(context.TODO(), filter, updates, upsertFlag)
	//jsonDataResult, err := JsonDataCollection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println("Data Insertion Error: ", err.Error())
		return err
	}
	fmt.Printf("Data Inserted: %v \tData Updated: %v\n", jsonDataResult.UpsertedCount, jsonDataResult.ModifiedCount)
	return nil
}

func FindData() error {
	err := ConnectDB()
	if err != nil {
		return err
	}
	var outputData []models.Data
	JsonDataCollection := Client.Database("testMongo").Collection("jsonData")
	outputDataCursor, err := JsonDataCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Collection Error: ", err.Error())
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = outputDataCursor.All(ctx, &outputData)
	if err != nil {
		fmt.Println("?? Error: ", err.Error())
		return err
	}
	for _, data := range outputData {
		fmt.Println(data.UserId, "\t", data.Id, "\t", data.Title)
	}
	return nil
}

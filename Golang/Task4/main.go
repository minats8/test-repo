/*
	Install Mongo and Setup Docker
	Interact with mongo with GO
	Make an Interface for file structure { mongo read and wirte }
	Add cmd line argument to decide the store  {file / db}
*/

package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const URL string = "https://jsonplaceholder.typicode.com/posts"
const connectionString string = "mongodb://localhost:27017"

func main() {

	fmt.Println("Enter the post you want to fetch: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	CheckError(err)
	//fetch post from json url
	dataFromPost := LoadPageData(strings.TrimSpace(input))
	//connection URI string format
	//mongodb://mongodb0.example.com:27017
	//configuring the client to use correct uri
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	CheckError(err)
	//timeout duration during which we try to connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//actual connecttion
	defer client.Disconnect(ctx)
	defer cancel()
	err = client.Connect(ctx)
	CheckError(err)
	//if connected properly we can ping the cluster from within our application
	err = client.Ping(ctx, readpref.Primary())
	CheckError(err)
	//printing list of database present

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	CheckError(err)
	fmt.Println("connection successful...", databases)

	jsonDataCollection := client.Database("testMongo").Collection("jsonData")
	/*
		Static data
			data := bson.D{
				{"userId", 1},
				{"id", 2},
				{"title", "sunt aut facere repellat provident occaecati excepturi optio reprehenderit"},
				{"body", "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"},
			}

	*/
	jsonDataResult, err := jsonDataCollection.InsertOne(context.TODO(), dataFromPost)
	CheckError(err)
	fmt.Println(jsonDataResult.InsertedID)

	//print all data from database
	var outputData []Data
	outputDataCursor, err := jsonDataCollection.Find(context.TODO(), bson.D{})
	CheckError(err)
	err = outputDataCursor.All(ctx, &outputData)
	CheckError(err)
	for index, data := range outputData {
		fmt.Println("\n", index, " ", data.UserId, " ", data.Id, " ", data.Title)
	}
}

func LoadPageData(postNO string) Data {
	newURL := URL + "/" + postNO
	resp, err := http.Get(newURL)
	CheckError(err)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	CheckError(err)
	var data Data
	err = json.Unmarshal(content, &data)
	CheckError(err)
	if data.Id == 0 {
		fmt.Println("Invalid post...")
		os.Exit(1)
	}
	return data
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
}

type Data struct {
	//bson:"<fieldname>,omitempty"
	//the omitempty means that if there is no data in the particular field,
	//when saved to MongoDB the field will not exist on the document rather than existing with an empty value.
	UserId int    `bson : "userId",omitempty`
	Id     int    `bson : "id",omitempty`
	Title  string `bson : "title",omitempty`
	Body   string `bson : "body",omitempty`
}

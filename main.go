package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"demo/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client mongo.Client

func InsertPost(title string, body string) {

	post := Post{title, body}

	collection := client.Database("game_report").Collection("info")

	insertResult, err := collection.InsertOne(context.TODO(), post)

	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("Inserted post with ID:", insertResult.InsertedID)

}

type Post struct {
	Title string `json:"title,omitempty"`

	Body string `json:"body,omitempty"`
}

func main() {
	
	myConfig := new(conf.Config)
	myConfig.InitConfig("config.ini")
	username := myConfig.Read("database","username")
	password := myConfig.Read("database","password")
	host     := myConfig.Read("database","host")
	db		 := myConfig.Read("database","db")

	connectstr := "mongodb+srv://" + username + ":" + password + "@" + host + "/" + db + "?retryWrites=true&w=majority"
	
    fmt.Println(connectstr)
	client1, err := mongo.NewClient(options.Client().ApplyURI(connectstr))
	if err != nil {
		log.Fatal(err)
	}
	client = *client1
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	InsertPost("中文", "测试内容")
}

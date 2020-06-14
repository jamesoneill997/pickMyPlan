package connection

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//set db address mongodb+srv://james:<password>@cluster0-cg8ph.mongodb.net/<dbname>?retryWrites=true&w=majority
var dbAddr = "mongodb+srv://james:Pickmyplan123@cluster0-cg8ph.mongodb.net/pickMyPlan?retryWrites=true&w=majority"

//Creates new client
var client, err = mongo.NewClient(options.Client().ApplyURI(dbAddr))
var ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)

//SetConnection connects to database
func SetConnection() *mongo.Client {
	//handle err
	if err != nil {
		fmt.Print("Error connecting to database\n\n")
		fmt.Println(err)
	}

	//Sets timeout for cancellation
	defer cancel()
	err = client.Connect(ctx)

	//returns pointer to mongo.Client type
	return client
}

//GetPort function to allow testing and production to run separately
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

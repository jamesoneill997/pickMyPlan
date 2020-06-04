package connection

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//set db address
var dbAddr = "mongodb://localhost:27017"

//Creates new client
var client, err = mongo.NewClient(options.Client().ApplyURI(dbAddr))
var ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)

//SetConnection connects to database (currently hosted on localhost)
func SetConnection() *mongo.Client {
	//handle err
	if err != nil {
		fmt.Println(err)
	}

	//Sets timeout for cancellation
	defer cancel()
	err = client.Connect(ctx)

	//returns pointer to mongo.Client type
	return client
}

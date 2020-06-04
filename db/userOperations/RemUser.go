package db

import (
	"context"
	"errors"

	connection "github.com/jamesoneill997/pickMyPlan/db/connection"
	"go.mongodb.org/mongo-driver/bson"
)

//Remove function connects to the db and removes user by unique username
func Remove(uName string) (int, error) {

	//setup db connection
	client := connection.SetConnection()
	userCol := connection.ConnectCollection(client, "users")

	//filter to find user
	filter := bson.D{
		bson.E{
			"username", uName,
		},
	}

	//search for user in mongodb and delete
	result, err := userCol.DeleteOne(context.TODO(), filter)

	//internal error or user does not exist
	if err != nil || result.DeletedCount == 0 {
		return -1, errors.New("User does not exist")
	}

	//0 reponse ok
	return 0, nil
}

package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

//UpdateDetails to update specified fields in users collection
func UpdateDetails(username string, field string, newValue interface{}) (int, error) {
	//setup connection to database
	client := SetConnection()
	conn := ConnectCollection(client, "users")

	//filter used to find user by unique username
	filter := bson.D{
		bson.E{
			"username", username,
		},
	}

	//desired update
	update := bson.D{
		{"$set", bson.M{field: newValue}},
	}

	//results of update, non-0 exit code success
	updateResult, err := conn.UpdateOne(context.TODO(), filter, update)

	//handle err
	if err != nil {
		return -1, errors.New("Invalid operation")
	} else if updateResult.MatchedCount == 0 {
		return -1, errors.New("Invalid username")
	}

	//Success, 0 exit status
	return 0, nil
}

package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

//UpdateDetails to update users
func UpdateDetails(username string, field string, newValue interface{}) {
	client := SetConnection()
	conn := ConnectCollection(client, "users")

	filter := bson.D{
		bson.E{
			"username", username,
		},
	}
	fmt.Println(field, newValue)
	update := bson.D{
		{"$set", bson.E{field, newValue}},
	}

	updateResult, err := conn.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		fmt.Println(err)
	} else if updateResult.MatchedCount == 0 {
		fmt.Println("Invalid username")
	}
}

package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func Remove(uName string) string {
	client := SetConnection()
	userCol := ConnectCollection(client, "users")
	filter := bson.D{
		bson.E{
			"username", uName,
		},
	}

	result, err := userCol.DeleteOne(context.TODO(), filter)

	if err != nil {
		return "Error"
	} else {
		fmt.Println(result)
		return "User successfully deleted"
	}
}

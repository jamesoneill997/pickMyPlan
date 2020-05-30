package db

import (
	"context"

	template "github.com/jamesoneill997/pickMyPlan/templates"
	"go.mongodb.org/mongo-driver/bson"
)

//FindUserByUsername function
func FindUserByUsername(username string) template.User {
	client := SetConnection()
	userCol := ConnectCollection(client, "users")
	user := template.User{}

	filter := bson.D{
		bson.E{
			"username", username,
		},
	}

	userCol.FindOne(context.TODO(), filter).Decode(&user)

	return user
}

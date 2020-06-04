package db

import (
	"context"
	"errors"

	connection "github.com/jamesoneill997/pickMyPlan/db/dbconnection"
	template "github.com/jamesoneill997/pickMyPlan/templates"
	"go.mongodb.org/mongo-driver/bson"
)

//FindUserByUsername function
func FindUserByUsername(username string) (template.User, error) {
	client := connection.SetConnection()
	userCol := connection.ConnectCollection(client, "users")

	//store results of find in user
	user := template.User{}

	//mongodb filter to find user by unique username
	filter := bson.D{
		bson.E{
			"username", username,
		},
	}

	//find user in db
	userCol.FindOne(context.TODO(), filter).Decode(&user)

	//if the user does not exist
	if len(user.Username) == 0 {
		return template.User{}, errors.New("User not found")
	}

	//success, user exists
	return user, nil
}

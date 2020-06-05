package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

//AddUser function adds a user to the DB, returns non-zero exit status for error
func AddUser(userCol mongo.Collection, user interface{}) int {
	//add user
	res, err := userCol.InsertOne(context.Background(), user)

	//handle err
	if err != nil {
		fmt.Println(err)
		return -1
	}

	//success, print object id
	fmt.Println(res.InsertedID)
	return 0

}

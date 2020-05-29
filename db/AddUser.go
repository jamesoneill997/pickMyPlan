package db

import (
	"context"
	"fmt"

	template "github.com/jamesoneill997/pickMyPlan/templates"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddUser(userCol mongo.Collection, user template.User) int {
	//add user
	res, err := userCol.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println(err)
		return -1
	} else {
		fmt.Println("User added successfully")
		fmt.Println(res)
		return 0
	}
}

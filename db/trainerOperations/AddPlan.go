package trainerdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/jamesoneill997/pickMyPlan/db/connection"
	"github.com/jamesoneill997/pickMyPlan/templates"
)

//AddPlan adds plan to plan collection
func AddPlan(plan templates.Program) (int, error) {
	conn := connection.SetConnection()
	coll := connection.ConnectCollection(conn, "plan")

	res, err := coll.InsertOne(context.Background(), plan)

	if err != nil {
		return -1, errors.New("Error when adding plan to db")
	}

	fmt.Println(res.InsertedID)
	return 0, nil
}

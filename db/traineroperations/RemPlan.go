package traineroperations

import (
	"context"
	"errors"
	"net/http"

	"github.com/jamesoneill997/pickMyPlan/api/user"
	"github.com/jamesoneill997/pickMyPlan/db/connection"
	"go.mongodb.org/mongo-driver/bson"
)

//RemPlan deletes a plan from the database
func RemPlan(w http.ResponseWriter, r *http.Request, name string) (int, error) {
	//preconditions - DB connection, collection connection and CurrUser
	conn := connection.SetConnection()
	coll := connection.ConnectCollection(conn, "Plan")
	trainer := user.CurrUser(w, r).Username

	//identify plan by name+creator
	filter := bson.D{
		bson.E{
			"name", name,
		},
		bson.E{
			"creator", trainer,
		},
	}

	//delete match from db
	res, err := coll.DeleteOne(context.TODO(), filter)

	//handle err
	if err != nil || res.DeletedCount == 0 {
		return -1, errors.New("Error removing user")
	}

	//success, 0 exit status
	return 0, nil
}

package traineroperations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jamesoneill997/pickMyPlan/api/user"
	"github.com/jamesoneill997/pickMyPlan/db/connection"
	"go.mongodb.org/mongo-driver/bson"
)

//RemPlan deletes a plan from the database
func RemPlan(w http.ResponseWriter, r *http.Request, name string) (int, error) {
	//preconditions - DB connection, collection connection and CurrUser
	conn := connection.SetConnection()
	coll := connection.ConnectCollection(conn, "plan")
	trainer := user.CurrUser(w, r).Username
	fmt.Println(name, trainer)

	//identify plan by name+creator
	filter := bson.M{"name": name, "creator": trainer}

	//delete match from db
	res, err := coll.DeleteOne(context.TODO(), filter)

	//handle err
	if err != nil || res.DeletedCount == 0 {
		return -1, err
	}

	//success, 0 exit status
	return 0, nil
}

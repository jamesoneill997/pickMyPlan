package trainer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jamesoneill997/pickMyPlan/api/user"
	"github.com/jamesoneill997/pickMyPlan/db/connection"
	"github.com/jamesoneill997/pickMyPlan/templates"
	"go.mongodb.org/mongo-driver/bson"
)

//ReadPlans reads all plans where creator==CurrUser()
func ReadPlans(w http.ResponseWriter, r *http.Request) {
	//setu db connection
	conn := connection.SetConnection()
	coll := connection.ConnectCollection(conn, "plan")

	//list to store plans
	planList := []templates.Program{}

	//current user's username
	trainer := user.CurrUser(w, r).Username

	//filter and search for all plans where creator == current user
	filter := bson.M{"creator": trainer}
	plans, findErr := coll.Find(context.TODO(), filter)

	//handle err
	if findErr != nil {
		fmt.Println(findErr)
		w.WriteHeader(503)
		w.Write([]byte("Internal server error"))
		return
	}

	//iterate through plan results
	for plans.Next(context.TODO()) {

		//template to store current find result
		p := templates.Program{}

		//decode and store current find result
		plans.Decode(&p)

		//append decoded result to plan list
		planList = append(planList, p)

	}

	//success
	w.Write([]byte("Plans succesfully fetched"))

	return
}

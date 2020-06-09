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
	conn := connection.SetConnection()
	coll := connection.ConnectCollection(conn, "plan")
	planList := []templates.Program{}

	//current user's username
	trainer := user.CurrUser(w, r).Username
	filter := bson.M{"creator": trainer}

	plans, findErr := coll.Find(context.TODO(), filter)

	if findErr != nil {
		fmt.Println(findErr)
		w.WriteHeader(503)
		w.Write([]byte("Internal server error"))
		return
	}

	for plans.Next(context.TODO()) {
		p := templates.Program{}
		plans.Decode(&p)
		planList = append(planList, p)

	}
	w.WriteHeader(200)
	w.Write([]byte("Plans succesfully fetched"))

	fmt.Println(planList)
	return
}

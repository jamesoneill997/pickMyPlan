package trainer

import (
	"encoding/json"
	"net/http"

	trainerdb "github.com/jamesoneill997/pickMyPlan/db/traineroperations"
	"github.com/jamesoneill997/pickMyPlan/templates"
)

//CreatePlan parses the request and adds a plan to the plan collection
func CreatePlan(w http.ResponseWriter, r *http.Request) {
	//handle post request
	switch r.Method {
	case "POST":

		//setup Program template and parse body
		plan := templates.Program{}
		decoder := json.NewDecoder(r.Body)
		decodErr := decoder.Decode(&plan)

		//handle err
		if decodErr != nil {
			w.WriteHeader(400)
			w.Write([]byte("Bad request"))
			return
		}

		//add plan to database plan collection
		res, err := trainerdb.AddPlan(plan)

		//handle err
		if res != 0 || err != nil {
			w.WriteHeader(503)
			w.Write([]byte("Internal server error"))
			return
		}

		//success - Creation response
		w.WriteHeader(201)
		w.Write([]byte("Plan successfully added"))

	//discard all other request methods
	default:
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
	}
}

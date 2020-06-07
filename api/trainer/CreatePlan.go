package trainer

import (
	"encoding/json"
	"net/http"

	"github.com/jamesoneill997/pickMyPlan/templates"
	trainerdb "github.com/jamesoneill997/pickmyplan/db/traineroperations"
)

func CreatePlan(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		plan := templates.Program
		decoder := json.NewDecoder(r.Body)
		decodErr := decoder.Decode(&plan)

		if decoder != nil {
			w.WriteHeader(400)
			w.Write([]byte("Bad request"))
			return
		}

		res, err := trainerdb.AddPlan(plan)

		if res != 0 || err != nil {
			w.WriteHeader(503)
			w.Write([]byte("Internal server error"))
			return
		}

		w.WriteHeader(201)
		w.Write([]byte("Plan successfully added"))
	default:
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
	}
}

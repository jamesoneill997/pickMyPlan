package trainer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jamesoneill997/pickMyPlan/db/traineroperations"
)

//DelPlan is a trainer operation to remove a plan from the DB
func DelPlan(w http.ResponseWriter, r *http.Request) {
	planName := make(map[string]string)
	decoder := json.NewDecoder(r.Body)
	decodErr := decoder.Decode(&planName)
	if decodErr != nil {
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}

	res, err := traineroperations.RemPlan(w, r, planName["Name"])

	if err != nil || res != 0 {
		fmt.Println(err, res)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Write([]byte("Plan successfully deleted"))
	return
}

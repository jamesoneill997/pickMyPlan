package trainer

import (
	"net/http"

	"github.com/jamesoneill997/pickMyPlan/db/traineroperations"
)

//DelPlan is a trainer operation to remove a plan from the DB
func DelPlan(w http.ResponseWriter, r *http.Request) {
	res, err := traineroperations.RemPlan(w, r, "Yoga with James")

	if err != nil || res != 0 {
		w.WriteHeader(503)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Plan successfully deleted"))
	return
}

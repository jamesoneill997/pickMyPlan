package user

import (
	"net/http"

	db "github.com/jamesoneill997/pickMyPlan/db"
)

//UpdateDetails func
func UpdateDetails(w http.ResponseWriter, r *http.Request) {
	query := GetAllQueries(w, r)
	field := query["field"][0]
	username := query["username"][0]
	newValue := query["value"][0]

	db.UpdateDetails(username, field, newValue)
}

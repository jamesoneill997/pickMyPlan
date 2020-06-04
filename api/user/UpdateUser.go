package user

import (
	"net/http"

	userOps "github.com/jamesoneill997/pickMyPlan/db/userOperations"
)

//UpdateDetails func
func UpdateDetails(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	//handle PATCH requests
	case "PATCH":
		//parse request
		query := GetAllQueries(w, r)
		field := query["field"][0]
		username := query["username"][0]
		newValue := query["value"][0]

		//update on the database
		result, err := userOps.UpdateDetails(username, field, newValue)

		//error with the database update
		if result != 0 || err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Internal server error"))
			return
		}

		//Success, 200 response
		w.WriteHeader(200)
		w.Write([]byte("Update successfully completed"))

	//ignore all other request methods
	default:
		w.WriteHeader(403)
		w.Write([]byte("Not authorized"))
		return
	}
}

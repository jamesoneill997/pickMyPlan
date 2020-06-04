package user

import (
	"net/http"
	"time"

	db "github.com/jamesoneill997/pickMyPlan/db/userOperations"
)

//Delete function removes user from database and expires active auth token
func Delete(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//Handle delete request
	case "DELETE":
		//find user to be deleted
		user := CurrUser(w, r)
		result, err := db.Remove(user.Username)

		//handle err
		if err != nil || result != 0 {
			w.WriteHeader(403)
			w.Write([]byte("Not Authorized"))
		} else {
			//set cookie expiry to now
			cookie := http.Cookie{Name: "Token", Expires: time.Now()}
			http.SetCookie(w, &cookie)

			//Successful, 200 status
			w.Write([]byte("User successfully deleted"))
		}

	default:
		//Ignore all other request methods
		w.WriteHeader(401)
		w.Write([]byte("Bad request"))
	}

}

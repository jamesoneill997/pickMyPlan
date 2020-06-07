package trainer

import (
	"net/http"

	"github.com/jamesoneill997/pickMyPlan/api/user"
)

//IsTrainer is middleware to verify account type for trainer-exclusive actions
func IsTrainer(endpoint func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	//returns http handler func to comply with requirements of http handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//gets current user
		currentUser := user.CurrUser(w, r)
		//ensure the current user is a trainer
		switch currentUser.Type {
		case "Trainer":
			endpoint(w, r)

		//discard anything else
		default:
			w.WriteHeader(403)
			w.Write([]byte("Unauthorized"))
			return
		}
	})

}

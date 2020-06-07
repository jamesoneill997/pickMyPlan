package trainer

import (
	"net/http"

	"github.com/jamesoneill997/pickMyPlan/api/user"
)

//IsTrainer is middleware to verify account type for trainer-exclusive actions
func IsTrainer(endpoint func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentUser := user.CurrUser(w, r)
		switch currentUser.Type {
		case "Trainer":
			endpoint(w, r)

		default:
			w.WriteHeader(403)
			w.Write([]byte("Unauthorized"))
			return
		}
	})

}

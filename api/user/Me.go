package user

import (
	"fmt"
	"net/http"

	db "github.com/jamesoneill997/pickMyPlan/db"
)

//GetMe will return user information which is only accessible to the current user
func GetMe(w http.ResponseWriter, r *http.Request) {

	//Parse value from key, value pair
	authToken := r.Header["Cookie"][0][6:]

	//Parse username from auth token
	username, err := GetUsername(authToken)

	//handle err
	if err != nil {
		w.WriteHeader(503)
		w.Write([]byte("Internal server error"))
		return
	}

	user, err := db.FindUserByUsername(username)

	if err != nil {
		w.WriteHeader(503)
		w.Write([]byte("Internal server error"))
		return
	}

	//Success, 200 response
	w.WriteHeader(200)
	w.Write([]byte("Profile fetched successfully"))
	fmt.Println(user)

}

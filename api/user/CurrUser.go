package user

import (
	"fmt"
	"net/http"

	db "github.com/jamesoneill997/pickMyPlan/db/userOperations"
	template "github.com/jamesoneill997/pickMyPlan/templates"
)

//CurrUser function is used to return current user of type template.User
func CurrUser(w http.ResponseWriter, r *http.Request) template.User {

	//Parse value from key, value pair
	if r.Header["Cookie"] == nil {
		return template.User{}
	}

	authToken := r.Header["Cookie"][0][6:]
	//Parse username from auth token
	username, err := GetUsername(authToken)

	//handle err
	if err != nil {
		w.WriteHeader(503)
		w.Write([]byte("Internal server error"))
		return template.User{}
	}

	user, err := db.FindUserByUsername(username)

	if err != nil {
		w.Write([]byte("Internal server error"))
		fmt.Println(err)
		return template.User{}
	}

	//Success

	w.Write([]byte("Profile fetched successfully\n"))
	return user
}

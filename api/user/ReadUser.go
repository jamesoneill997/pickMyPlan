package user

import (
	"fmt"
	"net/http"
	"net/url"

	db "github.com/jamesoneill997/pickMyPlan/db"
)

//GetQueryString parses a single query string (username) from a GET request
func GetQueryString(w http.ResponseWriter, r *http.Request) string {
	switch r.Method {
	//handle GET and DELETE requests
	case "GET", "DELETE":
		//parse username
		username, ok := r.URL.Query()["username"]

		//handle err
		if !ok || len(username[0]) == 0 {
			fmt.Println("Error with username")
			return "Error"
		}

		//success, return username
		return username[0]

	//ignore all other request methods
	default:
		return "Bad request"
	}
}

//GetAllQueries gets all query strings from a GET request
func GetAllQueries(w http.ResponseWriter, r *http.Request) url.Values {
	switch r.Method {

	//Handle GET, DELETE AND PATCH requests
	case "GET", "DELETE", "PATCH":
		query := r.URL.Query()
		return query

		//ignore all other request methods
	default:
		return nil
	}
}

//Read function reads user from the database using unique username
func Read(w http.ResponseWriter, r *http.Request) {

	//use above functions to parse string and find user
	uName := GetQueryString(w, r)
	user, findErr := db.FindUserByUsername(uName)

	//handle err
	if findErr != nil {
		w.WriteHeader(404)
		w.Write([]byte("Not found"))
		return
	}

	//Success, 200 response
	w.WriteHeader(200)
	w.Write([]byte(user.Username))
}

package user

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var t jwt.MapClaims

//GetMe will return user information which is only accessible to the current user
func GetMe(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header["Cookie"][0][6:]
	username, err := GetUsername(authToken)
	if err != nil {
		w.WriteHeader(503)
		w.Write([]byte("Internal server error"))
	}

	fmt.Println("Your username is: ", username)

}

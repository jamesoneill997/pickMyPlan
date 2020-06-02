package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

//needs to be env var in production environment
var mySigningKey = []byte("Myawesomesigningkeyisthisstring")

//IsAuthorized - Middleware function that checks if a user has acceess to an endpoint
func IsAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	//return http.Handler type
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		//handle GET requests
		case "GET":
			//check if cookie exists
			if r.Header["Cookie"] != nil {
				//get cookie from request header, 6: takes out "Token=" string which causes int64 error
				cookieBody := r.Header.Get("Cookie")[6:]
				//parse cookie body and check token, 2nd param is a function which returns the signingKey
				token, parseErr := jwt.Parse(cookieBody, func(token *jwt.Token) (interface{}, error) {
					//handle err
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, errors.New("Unauthorized")
					}

					//return signing key to parse
					return mySigningKey, nil
				})

				//handle err
				if parseErr != nil || !token.Valid {
					w.WriteHeader(403)
					w.Write([]byte("Unauthorized"))
					fmt.Println(parseErr)
					fmt.Printf("%T ----- %v", token, *token)
					return
				}

				//Authorized, 200 response and serve endpoint
				w.WriteHeader(200)
				w.Write([]byte("Access granted"))
				endpoint(w, r)
				return

			}
			//if no cookie, user does not have access
			w.WriteHeader(403)
			w.Write([]byte("Unauthorized"))
			return

		//ignore all other request methods
		default:
			w.WriteHeader(401)
			w.Write([]byte("Bad request"))
			return
		}
	})
}

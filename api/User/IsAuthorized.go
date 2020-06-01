package User

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("Myawesomesigningkeyisthisstring")

func IsAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Cookie"] != nil {
			cookieBody := r.Header.Get("Cookie")[6:]
			token, err := jwt.Parse(cookieBody, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}

				return mySigningKey, nil
			})

			if err != nil {
				fmt.Println("Error with signature", err, "\n", token)
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Println("Cookie is nil")
		}
	})
}

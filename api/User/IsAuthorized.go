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
			token, err := jwt.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTEwMTg0MTYsInVzZXIiOiJKYW1lcyBPJ05laWxsIn0.Oodjgbq3x8LR03js_4u6SnxTKAnaqdG9k8N3c3Vanpc", func(token *jwt.Token) (interface{}, error) {
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
			fmt.Println(r.Header)
			fmt.Println("Cookie is nil")
		}
	})
}

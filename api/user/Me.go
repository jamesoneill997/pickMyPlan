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

	token, err := jwt.ParseWithClaims(authToken, &t, func(token *jwt.Token) (interface{}, error) {
		return []byte("Myawesomesigningkeyisthisstring"), nil
	})

	if err != nil {
		fmt.Println(err)
	}
	claims := token.Claims.(*jwt.MapClaims)

	claimMap := make(map[string]interface{})

	for k, v := range *claims {
		claimMap[k] = v
	}

	fmt.Println(claimMap["user"])

}

package user

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)
var t jwt.MapClaims

//GetUsername function to parse username from object. Used to determine current user
func GetUsername(authToken string) (string, error) {

	//Parse token
	token, err := jwt.ParseWithClaims(authToken, &t, func(token *jwt.Token) (interface{}, error) {
		//call env var in production as opposed to signing key
		return []byte("Myawesomesigningkeyisthisstring"), nil
	})

	//handle err
	if err != nil {
		fmt.Println(err)
		return "Error", err
	}

	//Parse claims from token
	claims := token.Claims.(*jwt.MapClaims)

	//Cast claims to map for easy manipulation
	claimMap := make(map[string]interface{})

	//range over claims to get all
	for k, v := range *claims {
		claimMap[k] = v
	}

	//take user claim
	username := claimMap["user"].(string)

	//Success, return username
	return username, nil
}
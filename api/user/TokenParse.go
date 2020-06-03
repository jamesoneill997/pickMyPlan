package user

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func GetUsername(authToken string) (string, error) {
	token, err := jwt.ParseWithClaims(authToken, &t, func(token *jwt.Token) (interface{}, error) {
		return []byte("Myawesomesigningkeyisthisstring"), nil
	})

	if err != nil {
		fmt.Println(err)
		return "Error", err
	}
	claims := token.Claims.(*jwt.MapClaims)

	claimMap := make(map[string]interface{})

	for k, v := range *claims {
		claimMap[k] = v
	}

	username := claimMap["user"].(string)
	return string(username), nil
}

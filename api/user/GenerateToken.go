package user

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

//needs to be env var in production environment
var signingKey = []byte("Myawesomesigningkeyisthisstring")

//GenerateToken function will generate a jwt for the user. Returns token or err
func GenerateToken() (string, error) {
	//setup signing method
	token := jwt.New(jwt.SigningMethodHS256)

	//setup claims
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "James O'Neill"
	claims["exp"] = 2147483647

	//Sign with signing key
	tokenString, err := token.SignedString(signingKey)

	//handle err
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	//success, return token string
	return tokenString, nil
}

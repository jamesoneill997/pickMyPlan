package User

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

//needs to be env var in productioin environment
var signingKey = []byte("Myawesomesigningkeyisthisstring")

func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "James O'Neill"
	claims["exp"] = 2147483647

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}

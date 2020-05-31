package User

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	db "github.com/jamesoneill997/pickMyPlan/db"
	template "github.com/jamesoneill997/pickMyPlan/templates"
)

func GenerateToken(w http.ResponseWriter, r *http.Request, dbUser template.User) string {
	username := dbUser.Username
	jwtKey := []byte("my_secret_key")
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 10000,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, tokErr := token.SignedString(jwtKey)

	if tokErr != nil {
		return "Error signing token"
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenString,
	})

	currTokens := dbUser.Tokens

	currTokens = append(currTokens, tokenString)
	db.UpdateDetails(username, "Tokens", currTokens)

	return "Token generation successful"
}

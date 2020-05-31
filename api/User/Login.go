package User

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	db "github.com/jamesoneill997/pickMyPlan/db"
)

type creds struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

		decoder := json.NewDecoder(r.Body)
		user := creds{}
		decodeErr := decoder.Decode(&user)

		if decodeErr != nil {
			fmt.Println("decode error ", decodeErr)
			return
		}

		username := user.Username
		enteredPassword := user.Password

		dbUser := db.FindUserByUsername(username)

		result := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(enteredPassword))

		if result == nil {
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
				fmt.Println("Error signing token")
			}

			http.SetCookie(w, &http.Cookie{
				Name:  "token",
				Value: tokenString,
			})

			currTokens := dbUser.Tokens

			currTokens = append(currTokens, tokenString)
			db.UpdateDetails(username, "Tokens", currTokens)

		} else {
			fmt.Println("Unsuccessful")
		}

	default:
		return
	}
}

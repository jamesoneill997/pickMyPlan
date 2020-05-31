package User

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	db "github.com/jamesoneill997/pickMyPlan/db"
)

type creds struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
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
			fmt.Println("Successful login")
		} else {
			fmt.Println("Unsuccessful")
		}

	default:
		return
	}
}

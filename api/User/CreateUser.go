package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	db "github.com/jamesoneill997/pickMyPlan/db"
	template "github.com/jamesoneill997/pickMyPlan/templates"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	client := db.SetConnection()
	userCol := db.ConnectCollection(client, "users")
	switch r.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		u := template.User{}
		decodeErr := decoder.Decode(&u)
		password := []byte(u.Password)

		hashedPassword, passErr := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

		if decodeErr != nil || passErr != nil {
			return
		}
		u.Password = string(hashedPassword)
		db.AddUser(*userCol, u)
		userToken, err := GenerateToken()
		if err != nil {
			fmt.Println("Error generating token: ", err)
		}
		c := http.Cookie{Name: "Token", Value: userToken}

		http.SetCookie(w, &c)
		return
	default:
		return
	}
}

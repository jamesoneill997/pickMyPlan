package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	db "github.com/jamesoneill997/pickMyPlan/db"
	template "github.com/jamesoneill997/pickMyPlan/templates"
)

//CreateUser > add to db > generate auth token > set cookie
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//connect to database
	client := db.SetConnection()
	userCol := db.ConnectCollection(client, "users")

	//handle requests, POST is to be handled, all else should be rejected
	switch r.Method {
	case "POST":

		//read submitted data and store in user struct
		decoder := json.NewDecoder(r.Body)
		u := template.User{}
		decodeErr := decoder.Decode(&u)

		//handle err
		if decodeErr != nil {
			fmt.Println(decodeErr)
			w.WriteHeader(400)
			w.Write([]byte("Invalid data submitted"))
			return
		}

		//encrypt password
		password := []byte(u.Password)
		hashedPassword, passErr := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

		//handle err
		if passErr != nil {
			w.WriteHeader(500)
			w.Write([]byte("Internal Server Error"))
			return
		}

		//set user password to encrypted password before storage
		u.Password = string(hashedPassword)

		//add user to db
		create := db.AddUser(*userCol, u)

		//handle err
		if create != 0 {
			w.WriteHeader(500)
			w.Write([]byte("Error creating account. Please try again."))
			return
		}

		//generate auth token for user
		userToken, tokeErr := GenerateToken(u.Username)
		if tokeErr != nil {
			w.WriteHeader(500)
			w.Write([]byte("Internal Server Error"))
			return
		}

		//create and set cookie (no expiry)
		c := http.Cookie{Name: "Token", Value: userToken}

		http.SetCookie(w, &c)
		//creation response
		w.WriteHeader(201)
		w.Write([]byte("User created successfully"))
		return
	default:
		//Any request other than POST will be ignored
		w.WriteHeader(400)
		return
	}
}

package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"

	connection "github.com/jamesoneill997/pickMyPlan/db/connection"
	db "github.com/jamesoneill997/pickMyPlan/db/userOperations"

	template "github.com/jamesoneill997/pickMyPlan/templates"
)

//CreateUser > add to db > generate auth token > set cookie
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//connect to database
	client := connection.SetConnection()
	userCol := connection.ConnectCollection(client, "users")

	//handle requests, POST is to be handled, all else should be rejected
	switch r.Method {
	case "POST":

		//read submitted data and store in user struct
		decoder := json.NewDecoder(r.Body)

		data := map[string]interface{}{}

		decodeErr := decoder.Decode(&data)
		fmt.Println(data["Type"])
		switch data["Type"].(string) {
		case "Member":
			member := template.User{}
			decodeErr = mapstructure.Decode(data, &member)
			pw, err := HashPassword(member.Password)
			member.Password = pw
			if err != nil {
				w.WriteHeader(503)
				w.Write([]byte("Internal Server Error"))
			}
			add := db.AddUser(*userCol, member)
			if add != 0 {
				w.WriteHeader(503)
				w.Write([]byte("Internal Server Error"))
			}
		case "Trainer":
			trainer := template.Trainer{}
			decodeErr = mapstructure.Decode(data, &trainer)
			pw, err := HashPassword(trainer.Password)
			trainer.Password = pw

			if err != nil {
				w.WriteHeader(503)
				w.Write([]byte("Internal Server Error"))
			}
			add := db.AddUser(*userCol, trainer)
			if add != 0 {
				w.WriteHeader(503)
				w.Write([]byte("Internal Server Error"))
			}
		default:
			w.WriteHeader(400)
			w.Write([]byte("Bad request"))
		}

		//handle err
		if decodeErr != nil {
			fmt.Println(decodeErr)
			w.WriteHeader(400)
			w.Write([]byte("Invalid data submitted"))
			return
		}

		//generate auth token for user
		userToken, tokeErr := GenerateToken(data["Username"].(string))
		if tokeErr != nil {
			w.WriteHeader(500)
			w.Write([]byte("Internal Server Error"))
			return
		}

		//create and set cookie (no expiry)
		c := http.Cookie{Name: "Token", Value: userToken}

		http.SetCookie(w, &c)
		//creation response
		w.Write([]byte("User created successfully"))
		return
	default:
		//Any request other than POST will be ignored
		w.WriteHeader(400)
		return
	}
}

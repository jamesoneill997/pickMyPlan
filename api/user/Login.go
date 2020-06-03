package user

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	db "github.com/jamesoneill997/pickMyPlan/db"
)

//stores username password
type creds struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

//Login function reads username and password from request body, checks for user, logs in user and generates auth token
func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	//handle POST method
	case http.MethodPost:

		//store inputted data into creds struct
		decoder := json.NewDecoder(r.Body)
		user := creds{}
		decodeErr := decoder.Decode(&user)

		//handle err
		if decodeErr != nil {
			w.WriteHeader(500)
			w.Write([]byte("Internal Server Error"))
			return
		}

		//access user creds struct instance
		username := user.Username
		enteredPassword := user.Password

		//find user in the db by username
		dbUser, dbErr := db.FindUserByUsername(username)

		//handle err
		if dbErr != nil {
			w.WriteHeader(401)
			w.Write([]byte("Invalid credentials"))
			return
		}

		//encrypt entered password and compare with password on the db
		result := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(enteredPassword))

		//non zero response indicates no error
		if result == nil {
			//generate token and set cookie
			userToken, genErr := GenerateToken(username)
			c := http.Cookie{Name: "Token", Value: userToken}
			http.SetCookie(w, &c)

			//handle err
			if genErr != nil {
				w.WriteHeader(500)
				w.Write([]byte("Internal Server Error"))
				return
			}

			//Success, 200 response
			w.WriteHeader(200)
			w.Write([]byte("Login successful"))

		}

	//ignore all other request methods
	default:
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}
}

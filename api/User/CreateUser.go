package User

import (
	"encoding/json"
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
		return
	default:
		return
	}
}

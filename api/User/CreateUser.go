package User

import (
	"encoding/json"
	"net/http"

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
		err := decoder.Decode(&u)

		if err != nil {
			return
		}

		db.AddUser(*userCol, u)
		return
	default:
		return
	}
}

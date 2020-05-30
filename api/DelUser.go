package api

import (
	"fmt"
	"net/http"

	db "github.com/jamesoneill997/pickMyPlan/db"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		user := GetQueryString(w, r)
		fmt.Println(user)
		fmt.Println(db.Remove(user))

	}
}

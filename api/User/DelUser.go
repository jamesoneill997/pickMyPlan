package User

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/jamesoneill997/pickMyPlan/db"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		user := GetQueryString(w, r)
		fmt.Println(user)
		fmt.Println(db.Remove(user))
		cookie := http.Cookie{Name: "Token", Expires: time.Now()}
		http.SetCookie(w, &cookie)

	}
}

package user

import (
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		cookie := http.Cookie{Name: "Token", Expires: time.Now()}

		http.SetCookie(w, &cookie)
	}
}

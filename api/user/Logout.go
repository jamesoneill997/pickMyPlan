package user

import (
	"net/http"
	"time"
)

//Logout function instantly expires token
func Logout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//Handle POST method
	case "POST":
		//generate cookie and set
		cookie := http.Cookie{Name: "Token", Expires: time.Now()}
		http.SetCookie(w, &cookie)
		w.WriteHeader(200)
		w.Write([]byte("Successfully logged out"))

	//Ignore any other request method
	default:
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))

	}

}

package user

import (
	"net/http"
)

//GetMe will return user information which is only accessible to the current user
func GetMe(w http.ResponseWriter, r *http.Request) {
	me := CurrUser(w, r)

	w.Write([]byte(me.Username))
}

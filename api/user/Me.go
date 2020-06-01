package user

import (
	"fmt"
	"net/http"
)

//GetMe will return user information which is only accessible to the current user
func GetMe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are authorized to access this!")
}

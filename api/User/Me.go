package User

import (
	"fmt"
	"net/http"
)

func GetMe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are authorized to access this!")
}

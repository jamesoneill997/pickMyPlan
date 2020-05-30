package api

import (
	"fmt"
	"net/http"
)

func Read(w http.ResponseWriter, r *http.Request) string {
	switch r.Method {
	case http.MethodGet:
		username, ok := r.URL.Query()["username"]

		if !ok || len(username[0]) == 0 {
			fmt.Println("Error with username")
			return "Error"
		}

		return username[0]

	default:
		return "Bad Request"
	}
}

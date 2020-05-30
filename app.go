package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	api "github.com/jamesoneill997/pickMyPlan/api"
)

//server
var s = &http.Server{
	Addr:           ":8080",
	Handler:        nil,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 16,
}

func read(w http.ResponseWriter, r *http.Request) {
	//client := db.SetConnection()
	//userCol := db.ConnectCollection(client, "users")
	//user := template.User{}

	switch r.Method {
	case http.MethodGet:
		username, ok := r.URL.Query()["username"]

		if !ok || len(username[0]) == 0 {
			fmt.Println("Error with username")
			return
		}

		fmt.Println(username[0])

	default:
		return
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete profile\n")
}

//main function
func main() {
	fmt.Println("Server running on localhost:8080")
	http.HandleFunc("/create", api.CreateUser)
	http.HandleFunc("/user", read)
	http.HandleFunc("/delete", delete)

	log.Fatal(s.ListenAndServe())

}

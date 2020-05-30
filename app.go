package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	api "github.com/jamesoneill997/pickMyPlan/api"
	db "github.com/jamesoneill997/pickMyPlan/db"
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
	uName := api.GetQueryString(w, r)
	user := db.FindUserByUsername(uName)

	fmt.Println(user)
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

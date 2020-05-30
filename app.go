package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/jamesoneill997/pickMyPlan/db"
	template "github.com/jamesoneill997/pickMyPlan/templates"
)

//server
var s = &http.Server{
	Addr:           ":8080",
	Handler:        nil,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 16,
}

var client = db.SetConnection()

//handlers for endpoints
func create(w http.ResponseWriter, r *http.Request) {
	userCol := db.ConnectCollection(client, "users")
	switch r.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		u := template.User{}
		err := decoder.Decode(&u)

		if err != nil {
			return
		}

		db.AddUser(*userCol, u)
		return
	default:
		return
	}
}

func read(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Read profile\n")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete profile\n")
}

//main function
func main() {

	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)
	http.HandleFunc("/delete", delete)

	log.Fatal(s.ListenAndServe())

}

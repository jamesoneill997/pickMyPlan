package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	api "github.com/jamesoneill997/pickMyPlan/api"
	"github.com/jamesoneill997/pickMyPlan/db"
)

//server
var s = &http.Server{
	Addr:           ":8080",
	Handler:        nil,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 16,
}

func delete(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		user := api.GetQueryString(w, r)
		fmt.Println(user)
		fmt.Println(db.Remove(user))

	}
}

//main function
func main() {
	fmt.Println("Server running on localhost:8080")
	http.HandleFunc("/create", api.CreateUser)
	http.HandleFunc("/user", api.Read)
	http.HandleFunc("/delete", delete)

	log.Fatal(s.ListenAndServe())

}

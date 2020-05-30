package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	User "github.com/jamesoneill997/pickMyPlan/api/User"
)

//server
var s = &http.Server{
	Addr:           ":8080",
	Handler:        nil,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 16,
}

func main() {
	fmt.Println("Server running on localhost:8080")
	http.HandleFunc("/create", User.CreateUser)
	http.HandleFunc("/user", User.Read)
	http.HandleFunc("/delete", User.Delete)
	http.HandleFunc("/update", User.UpdateDetails)

	log.Fatal(s.ListenAndServe())

}

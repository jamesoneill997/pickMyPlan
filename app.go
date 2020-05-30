package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	api "github.com/jamesoneill997/pickMyPlan/api"
	db "github.com/jamesoneill997/pickMyPlan/db"
	template "github.com/jamesoneill997/pickMyPlan/templates"
	"go.mongodb.org/mongo-driver/bson"
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
	query := api.GetQueryString(w, r)

	client := db.SetConnection()
	userCol := db.ConnectCollection(client, "users")
	user := template.User{}

	filter := bson.D{
		bson.E{
			"username", query,
		},
	}

	userCol.FindOne(context.TODO(), filter).Decode(&user)

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

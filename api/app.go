package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//server
var s = &http.Server{
	Addr:           ":8080",
	Handler:        nil,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 16,
}

//handlers for endpoints
func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create profile\n")
}

func read(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Read profile\n")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete profile\n")
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)
	http.HandleFunc("/delete", delete)

	log.Fatal(s.ListenAndServe())
}

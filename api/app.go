package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/jamesoneill997/pickMyPlan/db"
	template "github.com/jamesoneill997/pickMyPlan/templates"

	"go.mongodb.org/mongo-driver/mongo"
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

func userConnection(client *mongo.Client) *mongo.Collection {
	//connect to users db
	collection := client.Database("pickMyPlan").Collection("users")
	return collection
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

//main function
func main() {
	userCol := userConnection(client)

	//test user
	u := template.User{
		Username:  "jamesoneill997",
		Gender:    "m",
		WeightKg:  100,
		HeightCm:  195,
		Build:     "muscular",
		Goals:     "bulk",
		Equipment: []string{"bodyweight", "resistance bands", "dumbells"},
		Statistics: map[string]int{
			"bench":    120,
			"Squat":    170,
			"Deadlift": 200,
		},
		Allergies:    nil,
		ProfileImage: "https://www.google.com/images",
		ProgressImages: []string{
			"https://www.google.com/images",
			"https://www.google.com/images",
		},
		PayAcctID: "1234",
	}

	db.AddUser(*userCol, u)

	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)
	http.HandleFunc("/delete", delete)

	log.Fatal(s.ListenAndServe())

}

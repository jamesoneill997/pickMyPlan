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
	switch r.Method {
	case http.MethodGet:
		fmt.Println(r.Method)
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		u := template.User{}
		err := decoder.Decode(&u)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(u)
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
	userCol := db.ConnectCollection(client, "users")

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

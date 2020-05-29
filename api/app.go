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

//user
type user struct {
	Username       string         `bson:"username" json:"username"`
	Gender         string         `bson:"gender" json:"username"`
	WeightKg       int            `bson:"weightKg" json:"username"`
	HeightCm       int            `bson:"heightCm" json:"username"`
	Build          string         `bson:"build" json:"username"`
	Goals          string         `bson:"goals" json:"username"`
	Equipment      []string       `bson:"equipment" json:"username"`
	Statistics     map[string]int `bson:"statistics" json:"username"`
	Allergies      []string       `bson:"allergies" json:"username"`
	ProfileImage   string         `bson:"profileImage" json:"username"`
	ProgressImages []string       `bson:"progressImages" json:"username"`
	PayAcctID      string         `bson:"payAcctID" json:"username"`
}

type trainer struct {
	username   string
	gender     string
	expertise  string
	experience string
	programs   string
	website    string
}

type program struct {
	category string
	exercies []struct {
		equipment  []string
		duration   int
		targetArea string
	}
	diet struct {
		meals []struct {
			ingredients []string
			allergies   []string
		}
	}
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
	//database connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	//test user
	u := user{
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

	//connect to users db
	collection := client.Database("pickMyPlan").Collection("users")
	//add user
	res, err := collection.InsertOne(context.Background(), u)

	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)
	http.HandleFunc("/delete", delete)

	log.Fatal(s.ListenAndServe())

	fmt.Println(res)
}

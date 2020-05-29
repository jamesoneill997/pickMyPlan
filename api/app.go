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

//database connection
func setConnection() *mongo.Client {
	var client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	return client
}

//user
type user struct {
	Username       string         `bson:"username" json:"username"`
	Gender         string         `bson:"gender" json:"gender"`
	WeightKg       int            `bson:"weightKg" json:"weightKg"`
	HeightCm       int            `bson:"heightCm" json:"heightCm"`
	Build          string         `bson:"build" json:"build"`
	Goals          string         `bson:"goals" json:"goals"`
	Equipment      []string       `bson:"equipment" json:"equipment"`
	Statistics     map[string]int `bson:"statistics" json:"statistics"`
	Allergies      []string       `bson:"allergies" json:"allergies"`
	ProfileImage   string         `bson:"profileImage" json:"profileImage"`
	ProgressImages []string       `bson:"progressImages" json:"progressImages"`
	PayAcctID      string         `bson:"payAcctID" json:"payAcctID"`
}

type trainer struct {
	Username   string   `bson:"username" json:"username"`
	Gender     string   `bson:"gender" json:"gender"`
	Expertise  []string `bson:"expertise" json:"expertise"`
	Experience string   `bson:"experience" json:"experience"`
	Programs   string   `bson:"programs" json:"programs"`
	Website    string   `bson:"website" json:"website"`
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
	client := setConnection()
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

	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)
	http.HandleFunc("/delete", delete)

	log.Fatal(s.ListenAndServe())

	fmt.Println(res)
}

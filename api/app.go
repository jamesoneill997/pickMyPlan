package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	template "github.com/jamesoneill997/pickMyPlan/templates"

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

func userConnection(client *mongo.Client) *mongo.Collection {
	//connect to users db
	collection := client.Database("pickMyPlan").Collection("users")
	return collection
}

func addUser(userCol mongo.Collection, user template.User) int {
	//add user
	res, err := userCol.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println(err)
		return -1
	} else {
		fmt.Println("User added successfully")
		fmt.Println(res)
		return 0
	}
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

	addUser(*userCol, u)

	http.HandleFunc("/create", create)
	http.HandleFunc("/read", read)
	http.HandleFunc("/delete", delete)

	log.Fatal(s.ListenAndServe())

}

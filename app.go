package main

import (
	"log"
	"net/http"
	"time"

	trainer "github.com/jamesoneill997/pickMyPlan/api/trainer"
	user "github.com/jamesoneill997/pickMyPlan/api/user"
)

//server
var s = &http.Server{
	Addr:           "https://pickmyplanapi.herokuapp.com",
	Handler:        nil,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 16,
}

func main() {
	//fmt.Println("Server running on localhost:8080")
	http.HandleFunc("/create", user.CreateUser)
	http.HandleFunc("/user", user.Read)
	http.HandleFunc("/delete", user.Delete)
	http.HandleFunc("/update", user.UpdateDetails)
	http.HandleFunc("/login", user.Login)
	http.HandleFunc("/logout", user.Logout)
	http.Handle("/me", user.IsAuthorized(user.GetMe))
	http.HandleFunc("/createplan", trainer.IsTrainer(trainer.CreatePlan))
	http.HandleFunc("/deleteplan", trainer.IsTrainer(trainer.DelPlan))
	http.HandleFunc("/readplans", trainer.IsTrainer(trainer.ReadPlans))

	log.Fatal(s.ListenAndServe())

}

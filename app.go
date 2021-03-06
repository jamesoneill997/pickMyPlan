package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	trainer "github.com/jamesoneill997/pickMyPlan/api/trainer"
	user "github.com/jamesoneill997/pickMyPlan/api/user"
	"github.com/jamesoneill997/pickMyPlan/db/connection"
)

//server
var s = &http.Server{
	Addr:           connection.GetPort(),
	Handler:        nil,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 16,
}

//main, handles endpoints
func main() {
	fmt.Println("Server running on " + connection.GetPort())
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

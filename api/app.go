package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//server
var s=&http.Server{
	Addr: ":8080",
	Handler: nil,
	ReadTimeout: 10 * time.Second,
	WriteTimeout: 10 * time.Second,
	MaxHeaderBytes: 1 << 16,
}


//handlers for endpoints
func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello\n")
}

func main() {
		
		http.HandleFunc("/hello", hello)
		log.Fatal(s.ListenAndServe())	
}
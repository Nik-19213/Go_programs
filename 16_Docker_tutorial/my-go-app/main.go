package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	fmt.Println("Welcome to Docker tutorial")

	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

//commands used
// docker build -t my-go-app .
//docker run -it -p 8080:8081 my-go-app    //container run until the session is running in bash
//docker run -d -p 8080:8081 my-go-app     //container run in background independently until we want
// docker ps 							   // to see all cotainers running in the machine
// docker kill <CONTAINER_ID>			   // to stop container execution

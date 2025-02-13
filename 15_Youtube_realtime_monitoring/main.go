// youtube-stats/main.go
package main

import (
	"fmt"
	"log"
	"monitoring/websocket"

	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "index.html")
}

// our new stats function which will expose any YouTube
// stats via a websocket connection
func stats(w http.ResponseWriter, r *http.Request) {
	// we call our new websocket package Upgrade
	// function in order to upgrade the connection
	// from a standard HTTP connection to a websocket one
	ws, err := websocket.HandleConnections(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	// we then call our Writer function
	// which continually polls and writes the results
	// to this websocket connection
	go websocket.Writer(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/ws", stats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("YouTube Subscriber Monitor")
	setupRoutes()

	// item, err := youtube.GetSubscribers()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v\n", item)
}

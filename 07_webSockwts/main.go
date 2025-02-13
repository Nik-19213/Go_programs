package main

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	// Handle new connections
	server.OnConnect("/", func(conn socketio.Conn) error {
		log.Println("New connection:", conn.ID())
		conn.Join("chat")
		return nil
	})

	// Handle chat messages
	server.OnEvent("/", "chat message", func(conn socketio.Conn, msg string) {
		log.Println("Message received from the client:", msg)
		server.BroadcastToRoom("/", "chat", "chat message", msg)
	})

	// Handle disconnects
	server.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		log.Println("Disconnected:", conn.ID(), "Reason:", reason)
	})

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Handle Socket.IO
	http.Handle("/socket.io/", server)

	// Start the server
	log.Println("Server is running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

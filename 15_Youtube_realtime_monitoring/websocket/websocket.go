// youtube-stats/websocket/websocket.go
package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"monitoring/youtube"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// We set our Read and Write buffer sizes
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections (you can change this for more security)
		return true
	},
}

// The Upgrade function will take in an incoming request and upgrade the request
// into a websocket connection
func HandleConnections(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// creates our websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ws, nil
}

func Writer(conn *websocket.Conn) {
	// we want to kick off a for loop that runs for the
	// duration of our websockets connection
	for {
		// we create a new ticker that ticks every 5 seconds
		ticker := time.NewTicker(5 * time.Second)
		// every time our ticker ticks
		for t := range ticker.C {
			// print out that we are updating the stats
			fmt.Printf("Updating Stats: %+v\n", t)
			// and retrieve the subscribers
			items, err := youtube.GetSubscribers()
			if err != nil {
				fmt.Println(err)
				return
			}
			// next we marshal our response into a JSON string
			jsonString, err := json.Marshal(items)
			if err != nil {
				fmt.Println(err)
				return
			}
			// and finally we write this JSON string to our WebSocket
			// connection and record any errors if there has been any
			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

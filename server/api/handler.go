package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// StreamHandler connects clients to the item stream via WebSocket.
func StreamHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unable to upgrade to websocket connection: %v", err)
		return
	}
	defer func() {
		_ = c.Close()
	}()

	for {
		msg := []byte("howdy pardner")
		if err = c.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Printf("failed to write message: %v", err)
			break
		}
		time.Sleep(1 * time.Second)
	}
}

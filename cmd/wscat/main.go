package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8000",
		Path:   "/",
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("failed to connect:", err)
	}
	defer func() {
		_ = c.Close()
	}()

	done := make(chan struct{})

	go func() {
		defer func() {
			_ = c.Close()
		}()
		defer close(done)
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("failed to read message:", err)
				return
			}
			log.Printf("received data: %s", string(msg))
		}
	}()

	for {
		select {
		case <-interrupt:
			log.Println("sending close message")
			closeMsg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
			err := c.WriteMessage(websocket.CloseMessage, closeMsg)
			if err != nil {
				log.Println("failed to close connection:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			_ = c.Close()
		}
	}
}

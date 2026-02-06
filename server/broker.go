package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Broker struct {
	// Events are pushed to this channel by the main events-gathering routine
	Notifier chan []byte

	// New client connections
	newClients chan chan []byte

	// Closed client connections
	closingClients chan chan []byte

	// Client connections registry
	clients map[chan []byte]bool
}

func NewBrokerServer() (broker *Broker) {
	broker = &Broker{
		Notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
	}

	go broker.listen()

	return
}

type Message struct {
	Name    string `json:"name"`
	Message string `json:"msg"`
}

func (broker *Broker) Stream(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Each connection registers its own message channel with the Broker's connections registry
	messageChan := make(chan []byte)

	// Signal the broker that we have a new connection
	broker.newClients <- messageChan

	// Remove this client from the map of connected clients when this handler exits
	defer func() {
		broker.closingClients <- messageChan
	}()

	// Listen to connection close using request context (replaces deprecated CloseNotifier)
	go func() {
		<-r.Context().Done()
		broker.closingClients <- messageChan
	}()

	for {
		// Write to the ResponseWriter - Server Sent Events compatible
		fmt.Fprintf(w, "data: %s\n\n", <-messageChan)

		// Flush the data immediately instead of buffering it for later
		flusher.Flush()
	}
}

func (broker *Broker) listen() {
	for {
		select {
		case s := <-broker.newClients:
			// A new client has connected - register their message channel
			broker.clients[s] = true

		case s := <-broker.closingClients:
			// A client has detached - stop sending them messages
			delete(broker.clients, s)

		case event := <-broker.Notifier:
			// We got a new event - send to all connected clients
			for clientMessageChan := range broker.clients {
				clientMessageChan <- event
			}
		}
	}
}

func (broker *Broker) Reload() {
	broker.Notifier <- []byte("1")
}

func (broker *Broker) BroadcastMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	_ = json.NewDecoder(r.Body).Decode(&msg)

	j, _ := json.Marshal(msg)
	broker.Notifier <- []byte(j)
	json.NewEncoder(w).Encode(msg)
}

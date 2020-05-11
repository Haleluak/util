package sse

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Broker struct {
	Notifier chan []byte
	newClients chan chan []byte
	closingClients chan chan []byte
	clients map[chan []byte] bool
}

func NewBroker() (broker * Broker)  {
	broker = &Broker{
		Notifier: make(chan []byte, 1),
		newClients: make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients: make(map[chan []byte] bool),
	}

	go broker.listen()
	return
}

func (b * Broker) ServeHTTP(rw http.ResponseWriter, req *http.Request)  {
	flusher, ok := rw.(http.Flusher)

	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	messageChan := make(chan []byte)
	b.newClients <- messageChan

	defer func() {
		b.closingClients <- messageChan
	}()

	notify := req.Context().Done()
	go func() {
		<-notify
		b.closingClients <- messageChan
	}()

	for {
		fmt.Fprintf(rw, "data: %s\n\n", <-messageChan)
		flusher.Flush()
	}

	log.Println("Finished HTTP request at ", req.URL.Path)
}

func (b * Broker) listen()  {
	for{
		select {
		case s:= <- b.newClients:
			b.clients[s] = true
			fmt.Println("Added new client")
		case s := <- b.closingClients:
			delete(b.clients, s)
			fmt.Print("Removed client %d registered clients")
		case event := <- b.Notifier:
			for clientMessageChan, _ := range b.clients {
				clientMessageChan <- event
			}
		}
	}
}

func run()  {
	broker := NewBroker()

	go func() {
		for{
			time.Sleep(time.Second * 2)
			eventString := fmt.Sprintf("the time is %v", time.Now())
			log.Println("Receiving event")
			broker.Notifier <- []byte(eventString)
		}
	}()

	log.Fatal("HTTP server error: ", http.ListenAndServe("localhost:3000", broker))
}
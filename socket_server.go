package main

import (
	"log"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

var server *gosocketio.Server

func initSocketServer() {
	server = gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Printf("Channel %s created", c.Id())
		c.Join("notifications")
	})
}

func broadcastUpdate(event string, data interface{}) {
	if server != nil {
		go server.BroadcastTo("notifications", event, data)
	}
}

func broadcastIndicator(indicator FileCopyInfo) {
	if server != nil {
		go server.BroadcastTo("notifications", "progress", indicator)
	}
}

package main

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

var server *socketio.Server

func initSocketServer() {
	var err error
	server, err = socketio.NewServer(nil)

	if err != nil {
		failOnError(err, "Impossible de créer le serveur")
	}

	server.OnConnect("/", func(c socketio.Conn) error {
		c.SetContext("")
		c.Join("notifications")

		log.Printf("Channel %s created", c.ID())
		return nil
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Printf("Channel %s encountered an error  ", s.ID(), e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Printf("Channel %s  Disconnected  ", s.ID(), reason)
	})
	go server.Serve()
}

func broadcastUpdate(event string, data interface{}) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Une erreur est survenue lors de l'envoi d'information WS", r)
		}
	}()
	if server != nil {
		log.Println(event, data)
		server.BroadcastToRoom("/", "notifications", event, data)
	}
}

func broadcastIndicator(indicator FileCopyInfo) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Une erreur est survenue lors de l'envoi d'information WS", r)
		}
	}()
	if server != nil {
		log.Println("progress", indicator)
		server.BroadcastToRoom("/", "notifications", "progress", indicator)
	}
}

package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var serveMux *http.ServeMux

func main() {

	prepareLogs()
	createRepository()
	log.Println("Dépôt créé et prêt")
	serveMux = http.NewServeMux()
	router = NewRouter()
	initMiddleware(router)

	RestartHTTPServer()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Une erreur est survenue après l'exposition du serveur", r)
			}
		}()
		log.Println("WebSocket Server online")
		serveMux.Handle("/socket.io/", server)
	}()

	RestartSocketServer()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, os.Kill)

	select {
	case <-sigChan:

		time.Sleep(5 * time.Second)
		logFile.Close()
		os.Exit(0)
	}
}

// RestartHTTPServer Redémarrage du serveur / module HTTP
func RestartHTTPServer() {

	initConfiguration()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Une erreur est survenue dans le code du Serveur HTTP", r)
			}
		}()

		log.Println("Serving at ", appConfig.httpListenURI(), "")
		log.Println(http.ListenAndServe(appConfig.httpListenURI(), router))
	}()

}

// RestartSocketServer Redémarrage du serveur / module de WebSocket
func RestartSocketServer() {

	initConfiguration()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Une erreur est survenue dans le code du Serveur HTTP", r)
			}
		}()

		log.Println("Serving Websocket at ", appConfig.socketListenURI(), "")
		log.Println(http.ListenAndServe(appConfig.socketListenURI(), serveMux))

	}()

}

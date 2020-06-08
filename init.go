package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/keybase/go-ps"
)

func main() {
	processes, err := ps.Processes()
	if err != nil {
		failOnError(err, "Error :")
	}
	procCount := 0
	for index := range processes {

		if strings.Contains(os.Args[0], processes[index].Executable()) {
			procCount++
		}

		if procCount > 1 {
			return
		}
	}
	prepareLogs()
	createRepository()
	log.Println("Dépôt créé et prêt")
	router = NewRouter()
	initMiddleware(router)
	initSocketServer()
	RestartHTTPServer()

	RestartSocketServer()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, os.Kill)

	select {
	case <-sigChan:
		server.Close()
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
		http.Handle("/socket.io/", server)
		log.Println("Serving Websocket at ", appConfig.socketListenURI(), "")
		log.Println(http.ListenAndServe(appConfig.socketListenURI(), nil))

	}()

}

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func initMiddleware(router *mux.Router) {

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s\t%s\t", r.Method, r.RequestURI)
			printRequest(r.RemoteAddr)
			constructHeaders(&w, r)
			next.ServeHTTP(w, r)

		})
	})

}

// LoggerHandler Middleware qui permet de logger les requêtes entrantes
func LoggerHandler(next http.HandlerFunc) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[BEGIN CALL] - %s\t%s\t", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)

		log.Printf(
			"[END CALL] - %s\t%s\t%s\t",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)

		w.WriteHeader(http.StatusOK)

	})
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("[ERROR] - %s: %s", msg, err)
	}
}

func printRequest(addr string) {
	log.Printf("[ %s ] - Request from %s ", time.Now().Format(time.RFC3339), addr)
}

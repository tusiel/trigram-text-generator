package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./config"
	"./middleware"
	"./routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan

		appCleanup()
		os.Exit(1)
	}()

	start()
}

func appCleanup() {
	log.Println("Shutting down server...")
}

func start() {
	router := mux.NewRouter()

	/** API Routes */
	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/learn", routes.HandleLearn).Methods("POST")
	api.HandleFunc("/generate", routes.HandleGenerate).Methods("GET")

	var handler http.Handler
	handler = router
	handler = handlers.LoggingHandler(os.Stdout, handler)
	handler = middleware.RemoveTrailingSlash(handler)

	srv := &http.Server{
		Handler:      handler,
		Addr:         config.GetString("listenAddress"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// "http://www.gutenberg.org/cache/epub/1661/pg1661.txt"

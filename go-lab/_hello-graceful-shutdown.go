package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var IsShuttingDown bool

func main() {
	// subscribe to SIGINT signals
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsShuttingDown {
			fmt.Fprintf(w, "We're Sorry, on shutting down process")
			return
		}
		fmt.Println("processing request...", r.URL.Path[1:])
		time.Sleep(5 * time.Second)
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
		fmt.Println("reply response...", r.URL.Path[1:])
	}))

	srv := &http.Server{Addr: ":8080", Handler: mux}

	go listenToSigTerm(quit, srv)

	log.Println("Server started at localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatalln("HTTPServer closed with error:", err)
		}
	}

	log.Println("Bye")
}

func listenToSigTerm(quit chan os.Signal, srv *http.Server) {
	<-quit

	log.Println("Shutting down server...")
	IsShuttingDown = true

	log.Println("Server gracefully stopped -with-no-shutdown-")
	time.Sleep(10 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("could not shutdown: %v", err)
	}
}

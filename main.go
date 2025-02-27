package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/product-api-microservice/handlers"
)

func main() {

	l := log.New(os.Stdout, "go-server: ", log.LstdFlags)

	hh := handlers.NewHello(l)
	hh2 := handlers.NewHeartbeat(l)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()

	sm.Handle("/hello", hh)
	sm.Handle("/", hh2)
	sm.Handle("/products", ph)

	s := &http.Server{
		Handler: sm,
		Addr: ":8000",
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func () {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Process for graceful shutdown

	sigChan := make(chan os.Signal, 2)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Println("Graceful Shutdown, due to :", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
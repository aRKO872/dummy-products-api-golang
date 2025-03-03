package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/product-api-microservice/handlers"
)

func main() {

	l := log.New(os.Stdout, "go-server: ", log.LstdFlags)

	hb := handlers.NewHeartbeat(l)
	pr := handlers.NewProducts(l)

	sm := mux.NewRouter()

	heartbeatRouter := sm.Methods(http.MethodGet).Subrouter()
	heartbeatRouter.HandleFunc("/heartbeat", hb.ServeHTTP)

	// sm.Handle("/products", ph)
	// sm.Handle("/products/{id}", ph)

	productGet := sm.Methods(http.MethodGet).Subrouter()
	productGet.HandleFunc("/products", pr.GetProducts)

	productPost := sm.Methods(http.MethodPost).Subrouter()
	productPost.HandleFunc("/products", pr.AddProduct)
	productPost.Use(pr.MiddlewareEncodingProduct)

	productPut := sm.Methods(http.MethodPut).Subrouter()
	productPut.HandleFunc("/products/{id:[0-9]+}", pr.UpdateProduct)
	productPut.Use(pr.MiddlewareEncodingProduct)

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

	sig := <-sigChan

	l.Println("Graceful Shutdown, due to :", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
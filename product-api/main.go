package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aRKO872/currency-grpc-service/protos/currency"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gohandler "github.com/gorilla/handlers"
	"github.com/product-api-microservice/handlers"
)

func main() {

	l := log.New(os.Stdout, "go-server: ", log.LstdFlags)

	hb := handlers.NewHeartbeat(l)

	sm := mux.NewRouter()

	gs, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.Fatal("err occured doing microservice call", err)
		l.Fatal(err)
	}

	cc := currency.NewCurrencyClient(gs)

	pr := handlers.NewProducts(l, &cc)

	// sm.Handle("/products", ph)
	// sm.Handle("/products/{id}", ph)

	productGet := sm.Methods(http.MethodGet).Subrouter()
	productGet.HandleFunc("/products", pr.GetProducts)
	productGet.HandleFunc("/product/{id}", pr.GetProductsSingle)
	productGet.HandleFunc("/heartbeat", hb.ServeHTTP)

	productPost := sm.Methods(http.MethodPost).Subrouter()
	productPost.HandleFunc("/products", pr.AddProduct)
	productPost.Use(pr.MiddlewareEncodingProduct)

	productPut := sm.Methods(http.MethodPut).Subrouter()
	productPut.HandleFunc("/products", pr.UpdateProduct)
	productPut.Use(pr.MiddlewareEncodingProduct)

	productDelete := sm.Methods(http.MethodDelete).Subrouter()
	productDelete.HandleFunc("/products/{id:[0-9]+}", pr.DeleteProduct)

	ops := middleware.RedocOpts{
		SpecURL: "/swagger.yaml",
	}

	sh := middleware.Redoc(ops, nil)
	productGet.Handle("/docs", sh)
	productGet.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	// We can add multiple referrers which can call this. Just adding one for the time being
	ch := gohandler.CORS(gohandler.AllowedOrigins([]string{"http://localhost:3000"}))

	// To allow everyone to access : 
	// ch := gohandler.CORS(gohandler.AllowedOrigins([]string{"*"}))

	s := &http.Server{
		Handler: ch(sm),			// CORS handler wrapping the router
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
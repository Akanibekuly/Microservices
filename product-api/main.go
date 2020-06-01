

package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handlers"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
	"golang.org/x/net/context"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//create the handlers
	ph := handlers.NewProducts(l)

	//create a new serve mux and register the handlers
	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	//create a new server
	s := &http.Server{
		Addr:         ":9090",           //configure the bind address
		Handler:      sm,                //set the default handler
		IdleTimeout:  120 * time.Second, //max time to read the request from client
		ReadTimeout:  1 * time.Second,   //max time to write response to the client
		WriteTimeout: 1 * time.Second,   //max time for connections using TCP Keep-Alive
	}

	//start the server
	go func() {
		log.Println("Starting server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	//block until a signal is recieved
	sig := <-sigChan
	l.Println("Recieved erminate, graceful shutdown", sig)

	//gracefully shutdown the server, waiting max 30 seconds for current operationsto complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}

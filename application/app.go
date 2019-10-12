package application

import (
	"context"
	"fmt"
	"github.com/crosscode-nl/incwebdemo/incrementer"
	"log"
	"net/http"
)

type App struct {
	incrementer incrementer.Incrementer
	done chan bool
}

func (a *App) handler(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "An error occurred.")
			a.done <- true
		}
	}()

	c := a.incrementer.Inc()
	fmt.Fprintf(w, "Counter: %v", c)
}


func (a *App) Run() {
	log.Println("Starting application on :5000")
	http.HandleFunc("/", a.handler)
	srv := &http.Server{Addr: ":5000"}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	<-a.done
	srv.Shutdown(context.TODO())
}

func New(incrementer incrementer.Incrementer) *App {
	return &App {
		incrementer: incrementer,
		done: make(chan bool),
	}
}

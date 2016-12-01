package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/samze/pratter/messages"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "port to start on")
	flag.Parse()

	startServer(port)
}

func startServer(port int) {
	address := fmt.Sprintf("127.0.0.1:%d", port)
	router := getRouter()

	addMessageHandlers(router)

	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func getRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
	})
	return r
}

func addMessageHandlers(r *mux.Router) {
	store := messages.NewMessageStore()
	addHandler := messages.AddHandler{&store}
	listHandler := messages.ListHandler{&store}
	r.Handle("/users/{user}/message", &addHandler).Methods("POST")
	r.Handle("/users/{user}/message", &listHandler).Methods("GET")
}

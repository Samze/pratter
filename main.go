package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var messages map[string][]message

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "port to start on")
	flag.Parse()

	messages = make(map[string][]message)
	startServer(port)
}

func startServer(port int) {
	address := fmt.Sprintf("127.0.0.1:%d", port)
	handler := getRouter()
	srv := &http.Server{
		Handler:      handler,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func getRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
	})
	r.HandleFunc("/users/{user}/message", postMessageHandler).Methods("POST")
	r.HandleFunc("/users/{user}/message", getMessageHandler).Methods("GET")

	return r
}

func postMessageHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	user := vars["user"]
	var msg message

	if req.Body == nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
	}

	res.WriteHeader(http.StatusCreated)
	if err := json.NewDecoder(req.Body).Decode(&msg); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	addMessage(user, msg)
}

func getMessageHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	user := vars["user"]

	userMessages := getMessages(user)

	if err := json.NewEncoder(res).Encode(&userMessages); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
}

func addMessage(user string, msg message) {
	messages[user] = append(messages[user], msg)
}

func getMessages(user string) []message {
	return messages[user]
}

type message struct {
	Text string `json: message`
}

package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("Hello World!!!")
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
}

func postTodoHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("Hello POST!!!")
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
}

func listTodosHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("Hello LIST!!!")
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	mux := http.NewServeMux()
	todoMux := http.NewServeMux()
	todoMux.HandleFunc("POST /", postTodoHandler)
	todoMux.HandleFunc("GET /", listTodosHandler)
	mux.Handle("/todos", todoMux)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

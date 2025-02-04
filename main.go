package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Todo struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var todoList []Todo

type PostTodoRequestBody struct {
	Text string `json:"text"`
}

func postTodoHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストボディの読み取り
	var requestBody PostTodoRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
		return
	}
	todoList = append(todoList, Todo{
		ID:   int64(len(todoList)),
		Text: requestBody.Text,
		Done: false,
	})
}

func listTodosHandler(w http.ResponseWriter, _ *http.Request) {
	b, err := json.Marshal(todoList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

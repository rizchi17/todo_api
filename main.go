package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strconv"
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

type UpdateTodoRequestBody struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
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

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
		return
	}
	// リクエストボディの読み取り
	var requestBody UpdateTodoRequestBody
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
		return
	}
	index := slices.IndexFunc(todoList, func(todo Todo) bool {
		return todo.ID == idInt64
	})
	if index == -1 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	todoList[index] = Todo{
		ID:   idInt64,
		Text: requestBody.Text,
		Done: requestBody.Done,
	}
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
		return
	}
	index := slices.IndexFunc(todoList, func(todo Todo) bool {
		return todo.ID == idInt64
	})
	if index == -1 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	todoList = slices.Delete(todoList, index, index+1)
}

func main() {
	todoMux := http.NewServeMux()
	todoMux.HandleFunc("POST /todos", postTodoHandler)
	todoMux.HandleFunc("GET /todos", listTodosHandler)
	todoMux.HandleFunc("PUT /todos/{id}", updateTodoHandler)
	todoMux.HandleFunc("DELETE /todos/{id}", deleteTodoHandler)

	err := http.ListenAndServe(":8080", todoMux)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type todo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []todo = make([]todo, 0)

func gettodosHundles(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		res, err := json.Marshal(todos)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "text/json")
		rw.Write(res)
		return
	}
	if r.Method == http.MethodPost {
		var body todo
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		todos = append(todos, body)
		rw.WriteHeader(http.StatusCreated)
		return
	}
}

func main() {

	// get all todos:GET /api/todos
	http.HandleFunc("/api/todos", gettodosHundles)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Lisntening on 0.0.0.0:8080")
}

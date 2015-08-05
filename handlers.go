package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	// 	var todo Todo
	// 	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if err := r.Body.Close(); err != nil {
	// 		panic(err)
	// 	}
	// 	if err := json.Unmarshal(body, &todo); err != nil {
	// 		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 		w.WriteHeader(422)
	// 		if err := json.NewEncoder(w).Encode(err); err != nil {
	// 			panic(err)
	// 		}
	// 	}

	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(http.StatusCreated)
	// 	if err := json.NewEncoder(w).Encode(t); err != nil {
	// 		panic(err)
	// 	}
}

func TodoEdit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

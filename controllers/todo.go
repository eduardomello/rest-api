package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eduardomello/rest-api/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TodoController struct {
	C *mgo.Collection
}

func NewTodoController(s *mgo.Session) *TodoController {
	tc := TodoController{s.DB("rest_api").C("todos")}
	return &tc
}

func (tc TodoController) GetTodo(w http.ResponseWriter, r *http.Request) {
	oid := GetId(mux.Vars(r))
	if oid == "" {
		w.WriteHeader(404)
		return
	}

	todo := models.Todo{}

	if err := tc.C.FindId(oid).One(&todo); err != nil {
		w.WriteHeader(404)
		return
	}

	todo_json, _ := json.Marshal(todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", todo_json)
}

func (tc TodoController) ListTodos(w http.ResponseWriter, r *http.Request) {
	var results []bson.M

	if err := tc.C.Find(nil).All(&results); err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintln(w, "[")
	for i, obj := range results {
		if i > 0 {
			fmt.Fprintln(w, ",")
		}
		todo_json, _ := json.Marshal(obj)
		fmt.Fprintf(w, "%s", todo_json)
	}
	fmt.Fprintln(w, "]")

}

func (tc TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}

	json.NewDecoder(r.Body).Decode(&todo)

	todo.Id = bson.NewObjectId()

	tc.C.Insert(todo)

	todo_json, _ := json.Marshal(todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", todo_json)
}

func (tc TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	oid := GetId(mux.Vars(r))
	if oid == "" {
		w.WriteHeader(404)
		return
	}

	if err := tc.C.RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)

}

func (tc TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	oid := GetId(mux.Vars(r))
	if oid == "" {
		w.WriteHeader(404)
		return
	}

	todo := models.Todo{}

	json.NewDecoder(r.Body).Decode(&todo)

	if err := tc.C.UpdateId(oid, todo); err != nil {
		w.WriteHeader(400)
		fmt.Println(err)
		fmt.Println(oid)
		return
	}

	if err := tc.C.FindId(oid).One(&todo); err != nil {
		w.WriteHeader(404)
		return
	}

	todo_json, _ := json.Marshal(todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", todo_json)

}

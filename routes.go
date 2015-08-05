package main

import (
	"net/http"

	"github.com/eduardomello/rest-api/controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var tc = controllers.NewTodoController(getSession())

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ListTodos",
		"GET",
		"/todos",
		tc.ListTodos,
	},
	Route{
		"GetTodo",
		"GET",
		"/todos/{todoId}",
		tc.GetTodo,
	},
	Route{
		"CreateTodo",
		"POST",
		"/todos",
		tc.CreateTodo,
	},
	Route{
		"UpdateTodo",
		"PUT",
		"/todos/{todoId}",
		tc.UpdateTodo,
	},
	Route{
		"DeleteTodo",
		"DELETE",
		"/todos/{todoId}",
		tc.DeleteTodo,
	},
}

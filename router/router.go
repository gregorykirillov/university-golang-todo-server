package router

import (
	"Go-todo-list/routes"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", routes.GetAllTasks).Methods("GET")
	router.HandleFunc("/task/{id}", routes.GetTaskByID).Methods("GET")
	router.HandleFunc("/addTask", routes.AddTask).Methods("POST")
	router.HandleFunc("/updateTask/{id}", routes.UpdateTask).Methods("PUT")
	router.HandleFunc("/deleteTask/{id}", routes.DeleteTask).Methods("DELETE")
}

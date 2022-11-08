package router

import (
	"net/http"
	"todo-app/routes"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	router.Use(mux.CORSMethodMiddleware(router))
	router.HandleFunc("/tasks", routes.GetAllTasks).Methods(http.MethodGet)
	router.HandleFunc("/task/{id}", routes.GetTaskByID).Methods(http.MethodGet)
	router.HandleFunc("/addTask", routes.AddTask).Methods(http.MethodPost)
	router.HandleFunc("/updateTask/{id}", routes.UpdateTask).Methods(http.MethodPut)
	router.HandleFunc("/deleteTask/{id}", routes.DeleteTask).Methods(http.MethodDelete)
}

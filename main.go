package main

import (
	"net/http"

	"Go-todo-list/initializers"
	"Go-todo-list/models"
	"Go-todo-list/router"

	"github.com/gorilla/mux"
)

func Initialize() {
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Task{})
}
func main() {
	Initialize()
	muxRouter := mux.NewRouter()

	router.InitRoutes(muxRouter)

	http.ListenAndServe(":5000", muxRouter)
}

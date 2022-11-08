package main

import (
	"net/http"

	"todo-app/initializers"
	"todo-app/models"
	"todo-app/router"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Initialize() {
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Task{})
}
func main() {
	Initialize()
	muxRouter := mux.NewRouter()

	router.InitRoutes(muxRouter)

	corsSettings := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	})

	handler := corsSettings.Handler(muxRouter)

	http.ListenAndServe(":5000", handler)
}

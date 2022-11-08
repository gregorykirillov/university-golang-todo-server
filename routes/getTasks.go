package routes

import (
	"encoding/json"
	"net/http"

	"todo-app/initializers"
	"todo-app/models"

	"github.com/gorilla/mux"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var tasks []models.Task
	result := initializers.DB.Find(&tasks)

	if result.Error != nil {
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task

	params := mux.Vars(r)
	result := initializers.DB.First(&task, params["id"])

	if result.Error != nil {
		return
	}

	json.NewEncoder(w).Encode(task)
}

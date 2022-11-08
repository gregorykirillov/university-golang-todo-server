package routes

import (
	"encoding/json"
	"net/http"
	"todo-app/initializers"
	"todo-app/models"

	"github.com/gorilla/mux"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task
	params := mux.Vars(r)
	result := initializers.DB.First(&task, params["id"])

	if result.Error != nil {
		return
	}

	json.NewDecoder(r.Body).Decode(&task)
	initializers.DB.Save(&task)

	json.NewEncoder(w).Encode(task)
}

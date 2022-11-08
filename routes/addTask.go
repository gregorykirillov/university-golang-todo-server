package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"todo-app/initializers"
	"todo-app/models"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	dbErr := initializers.DB.Create(&task)
	if dbErr.Error != nil {
		log.Fatalf(dbErr.Error.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(task)
}

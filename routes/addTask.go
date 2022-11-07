package routes

import (
	"Go-todo-list/initializers"
	"Go-todo-list/models"
	"encoding/json"
	"log"
	"net/http"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	if dbErr := initializers.DB.Create(&task); dbErr.Error != nil {
		log.Fatalf(dbErr.Error.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := make(map[string]string)
	resp["status"] = "ok"

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

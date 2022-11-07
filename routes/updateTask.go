package routes

import (
	"encoding/json"
	"net/http"

	"Go-todo-list/initializers"
	"Go-todo-list/models"

	"github.com/gorilla/mux"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	params := mux.Vars(r)

	if result := initializers.DB.First(&task, params["id"]); result != nil {
		resp := make(map[string]string)
		resp["status"] = "false"
		resp["message"] = result.Error.Error()

		jsonResp, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)

		return
	}

	json.NewDecoder(r.Body).Decode(&task)
	initializers.DB.Save(&task)

	json.NewEncoder(w).Encode(task)
}

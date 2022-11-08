package routes

import (
	"encoding/json"
	"net/http"
	"todo-app/initializers"
	"todo-app/models"

	"github.com/gorilla/mux"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result := initializers.DB.Delete(&models.Task{}, params)

	if result.Error != nil {
		resp := make(map[string]string)
		resp["status"] = "false"
		resp["message"] = result.Error.Error()

		jsonResp, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)

		return
	}

	json.NewEncoder(w).Encode(params)
}

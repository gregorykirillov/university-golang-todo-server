package routes

import (
	"Go-todo-list/initializers"
	"Go-todo-list/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	println("Herere")
	params := mux.Vars(r)
	if result := initializers.DB.Delete(&models.Task{}, params); result.Error != nil {
		resp := make(map[string]string)
		resp["status"] = "false"
		resp["message"] = result.Error.Error()

		jsonResp, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)

		return
	}
}

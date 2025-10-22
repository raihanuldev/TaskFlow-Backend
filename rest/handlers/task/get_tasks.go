package task

import (
	"encoding/json"
	"net/http"
	database "taskflow/Database"
)

func HandleGetTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != "GET" {
		//do error through
		http.Error(w, "Please Sent GET Request", 400)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(database.Stores)
}

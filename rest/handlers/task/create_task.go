package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	database "taskflow/Database"
)


func HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		//do error through
		http.Error(w, "Please Sent POST Request", 400)
		return
	}
	var newTask database.Task
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newTask)

	if err != nil {
		fmt.Println("Error")
		http.Error(w, "Please give  me a Vaild json", 400)
		return
	}
	newTask.ID = len(database.Stores) + 1
	database.Stores = append(database.Stores, newTask)

	// return api response
	encode := json.NewEncoder(w)
	encode.Encode(newTask)

}

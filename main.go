package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Task Struct

type Task struct {
	ID       int    `json:"id"`
	TaskName string `json:"task_name"`
	Details  string `json:"details"`
	Status   string `json:"status"`
}

var taskList []Task

// Handle Add Task

func HandleGetTask(w http.ResponseWriter, r *http.Request) {

	// Header Config set
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	// Request Method Catch
	if r.Method != "GET" {
		//do error through
		http.Error(w, "Please Sent GET Request", 400)
		return
	}
	//create a encoder for senting response for Tasks
	encoder := json.NewEncoder(w)
	encoder.Encode(taskList)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/task", HandleGetTask)

	//Server Config
	fmt.Println("Server is running on port, 3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
	}
}

func init() {
	taskList = append(taskList,
		Task{
			ID:       1,
			TaskName: "Learn Go Basics",
			Details:  "Understand variables, loops, and functions",
			Status:   "todo",
		},
		Task{
			ID:       2,
			TaskName: "Build REST API",
			Details:  "Practice with net/http and JSON handling",
			Status:   "in-progress",
		},
		Task{
			ID:       3,
			TaskName: "Read Go Concurrency",
			Details:  "Learn goroutines and channels",
			Status:   "pending",
		},
		Task{
			ID:       4,
			TaskName: "Implement Task Update Feature",
			Details:  "Add update and delete logic to task API",
			Status:   "done",
		},
	)
}

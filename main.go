package main

import (
	"fmt"
	"net/http"
	"taskflow/rest/handlers/task"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/task", task.HandleGetTask)
	mux.HandleFunc("/api/v1/create-task", task.HandleCreateTask)

	//Server Config
	fmt.Println("Server is running on port, 3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
	}
}

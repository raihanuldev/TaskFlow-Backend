package main

import (
	"fmt"
	"net/http"
)

// Task Struct

type Task struct{
	Id int
	TaskName string
	
}



// Handle Add Task

func HandleAddTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/addTask", HandleAddTask)


	//Server Config
	fmt.Println("Server is running on port, 3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
	}
}

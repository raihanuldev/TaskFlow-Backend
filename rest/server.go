package rest

import (
	"fmt"
	"net/http"
	"os"
	"taskflow/rest/handlers/task"
)

type Server struct {
	taskHandler task.Handler
}

// step 2
func NewServer(taskHandler *task.Handler) *Server {
	return &Server{
		taskHandler: *taskHandler,
	}
}

func (server *Server) Start() {

	mux := http.NewServeMux()

	server.taskHandler.RegisterRoutes(mux)

	fmt.Println("Server is running on port, ", 3000)
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
		os.Exit(1)
	}
}

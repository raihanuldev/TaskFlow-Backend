package cmd

import (
	"taskflow/rest"
	"taskflow/rest/handlers/task"
)

func Serve() {
	taskHandler := task.NewHandler()

	// server := rest.NewServer(productHandler, userHandler, *cnf)
	server:= rest.NewServer(taskHandler)
	server.Start()
}

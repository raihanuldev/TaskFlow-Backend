package task

import (
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("GET /api/v1/task", http.HandlerFunc(HandleGetTask))
	mux.Handle("POST /api/v1/create-task", http.HandlerFunc(HandleCreateTask))

}

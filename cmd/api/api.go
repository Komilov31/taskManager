package api

import (
	"log"
	"net/http"

	"github.com/Komilov31/TaskManagerApi/internal/handler"
	"github.com/Komilov31/TaskManagerApi/internal/repository"
	"github.com/Komilov31/TaskManagerApi/internal/service"
	"github.com/gorilla/mux"
)

type Server struct {
	address string
}

func NewServer(address string) *Server {
	return &Server{address: address}
}

func (s *Server) Run() error {
	router := mux.NewRouter()

	repository := repository.NewRepository()
	service := service.NewTaskManager(repository)
	handler := handler.NewHandler(service)
	handler.RegisterRoutes(router)

	log.Println("Listening on", s.address)
	return http.ListenAndServe(s.address, router)
}

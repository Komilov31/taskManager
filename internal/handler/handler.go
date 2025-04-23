package handler

import (
	"errors"
	"net/http"
	"strconv"

	pkg "github.com/Komilov31/TaskManagerApi/pkg/utills"
	"github.com/gorilla/mux"
)

type Handler struct {
	taskService TaskService
}

func NewHandler(service TaskService) *Handler {
	return &Handler{taskService: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/status", h.GetStatusHandler).Methods(http.MethodGet)
	router.HandleFunc("/result", h.GetResultHandler).Methods(http.MethodGet)
	router.HandleFunc("/newtask", h.PostTaskHandler).Methods(http.MethodPost)
}

func (h *Handler) GetStatusHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	taskId, err := strconv.Atoi(id)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, errors.New("sent id is not number"))
		return
	}

	status, err := h.taskService.GetTaskStatus(taskId)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err)
		return
	}

	pkg.WriteJson(w, http.StatusOK, map[string]any{"id": taskId, "status": status})
}

func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	task := h.taskService.StartTask()
	pkg.WriteJson(w, http.StatusOK, task)
}

func (h *Handler) GetResultHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	taskId, err := strconv.Atoi(id)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, errors.New("sent id is not number"))
		return
	}

	result, err := h.taskService.GetTaskResult(taskId)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err)
		return
	}

	pkg.WriteJson(w, http.StatusOK, map[string]any{"id": taskId, "result": result})
}

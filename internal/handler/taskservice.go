package handler

import pkg "github.com/Komilov31/TaskManagerApi/pkg/types"

type TaskService interface {
	StartTask() *pkg.Task
	GetTaskStatus(id int) (string, error)
	GetTaskResult(id int) (string, error)
}

package service

import (
	"log"
	"math/rand"
	"time"

	pkg "github.com/Komilov31/TaskManagerApi/pkg/types"
)

type TaskManagerService struct {
	store Store
}

func NewTaskManager(store Store) *TaskManagerService {
	return &TaskManagerService{store: store}
}

func (t *TaskManagerService) GetTaskStatus(id int) (string, error) {
	taskStatus, err := t.store.GetTaskStatus(id)
	if err != nil {
		return "", err
	}

	return taskStatus, err
}

func (t *TaskManagerService) StartTask() *pkg.Task {

	task := new(pkg.Task)
	task.Status = "pending"
	task.Result = "not ready yet"

	taskId := t.store.StoreTask(task)
	task.ID = taskId

	go func() {
		log.Println("Task with id =", taskId, "started")
		time.Sleep(time.Duration((rand.Int63n(3) + 3) * int64(time.Minute)))
		task.Status = "finished"
		task.Result = "this is result of your task"
		log.Println("Task with id = ", taskId, "finished")
	}()

	return task
}

func (t *TaskManagerService) GetTaskResult(id int) (string, error) {
	task, err := t.store.GetTask(id)
	if err != nil {
		return "", err
	}

	return task.Result, err
}

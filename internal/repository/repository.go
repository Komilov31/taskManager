package repository

import (
	"errors"
	"sync"

	pkg "github.com/Komilov31/TaskManagerApi/pkg/types"
)

type Repository struct {
	storage   map[int]*pkg.Task
	currentId int
	mu        sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{
		storage: map[int]*pkg.Task{},
		mu:      sync.RWMutex{},
	}
}

func (r *Repository) GetTaskStatus(id int) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.storage[id]
	if !ok {
		return "", errors.New("task with provided id does not exist")
	}

	return task.Status, nil
}

func (r *Repository) GetTask(id int) (*pkg.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.storage[id]
	if !ok {
		return &pkg.Task{}, errors.New("task with provided id does not exist")
	}

	return task, nil
}

func (r *Repository) StoreTask(task *pkg.Task) int {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.storage[r.currentId] = task
	r.currentId++

	return r.currentId
}

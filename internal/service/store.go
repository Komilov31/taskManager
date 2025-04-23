package service

import pkg "github.com/Komilov31/TaskManagerApi/pkg/types"

type Store interface {
	GetTaskStatus(int) (string, error)
	StoreTask(*pkg.Task) int
	GetTask(int) (*pkg.Task, error)
}

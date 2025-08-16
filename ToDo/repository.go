package qtodo

import (
	"fmt"
	"sync"
)

type Database interface {
	GetTaskList() []Task
	GetTask(string) (Task, error)
	SaveTask(Task) error
	DelTask(string) error
}

type database struct {
	taskList []Task
	mu sync.RWMutex
}

func (db *database) GetTaskList() []Task {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.taskList
}

func (db *database) GetTask(name string) (Task, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	for _, task := range db.taskList {
		if task.GetName() == name {
			return task, nil
		}
	}
	return nil, fmt.Errorf("task not found")
}

func (db *database) SaveTask(task Task) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	for _, t := range db.taskList {
		if t.GetName() == task.GetName() {
			return fmt.Errorf("task already exists")
		}
	}
	db.taskList = append(db.taskList, task)
	return nil
}

func (db *database) DelTask(name string) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	for i, task := range db.taskList {
		if task.GetName() == name {
			db.taskList = append(db.taskList[:i], db.taskList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

func NewDatabase() *database {
	return &database{
		taskList: make([]Task, 0),
		mu: sync.RWMutex{},
	}
}
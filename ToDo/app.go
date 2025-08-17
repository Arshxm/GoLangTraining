package qtodo

import (
	"time"
)

type App interface {
	StartTask(string) error
	StopTask(string)
	AddTask(string, string, time.Time, func(), bool) error
	DelTask(string) error
	GetTaskList() []Task
	GetTask(string) (Task, error)
}

type app struct {
	taskList map[string]Task
	dataDir  Database
	timers   map[string]*time.Timer
}

func (a *app) AddTask(name, description string, dueTime time.Time, callback func(), isTemp bool) error {
	task, err := NewTaskWithStopAction(callback, nil, dueTime, name, description, isTemp)
	if err != nil {
		return err
	}

	a.taskList[name] = task
	err = a.dataDir.SaveTask(task)
	if err != nil {
		delete(a.taskList, name)
		return err
	}
	return nil
}

func (a *app) DelTask(name string) error {
	err := a.dataDir.DelTask(name)
	if err != nil {
		return err
	}
	delete(a.taskList, name)
	return nil
}

func (a *app) GetTaskList() []Task {
	return a.dataDir.GetTaskList()
}

func (a *app) GetTask(name string) (Task, error) {
	task, err := a.dataDir.GetTask(name)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (a *app) StartTask(name string) error {
	task, err := a.dataDir.GetTask(name)
	if err != nil {
		return err
	}

	duration := task.GetAlarmTime().Sub(time.Now())

	if duration <= 0 {
		task.DoAction()
		if task.IsTemp() {
			a.DelTask(name)
		}
		return nil
	}
	timer := time.AfterFunc(duration, func() {
		task.DoAction()
		if task.IsTemp() {
			a.DelTask(name)
		}
		delete(a.timers, name)
	})

	a.timers[name] = timer

	return nil
}

func (a *app) StopTask(name string) {
	if timer, exists := a.timers[name]; exists {
		timer.Stop()
		delete(a.timers, name)
	}
}

func NewApp(db Database) *app {
	return &app{
		taskList: make(map[string]Task),
		dataDir:  db,
		timers:   make(map[string]*time.Timer),
	}
}

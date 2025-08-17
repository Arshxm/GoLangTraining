package qtodo

import (
	"errors"
	"time"
)

type TaskState int

const (
	TaskPending TaskState = iota
	TaskRunning
	TaskStopped
	TaskCompleted
)

type Task interface {
	DoAction()
	StopAction()
	GetAlarmTime() time.Time
	GetAction() func()
	GetName() string
	GetDescription() string
	IsTemp() bool
	GetState() TaskState
	SetState(TaskState)
}

type task struct {
	action      func()
	stopAction  func()
	alarmTime   time.Time
	name        string
	description string
	isTemp      bool
	state       TaskState
}

func (t *task) DoAction() {
	if t.state != TaskStopped {
		t.state = TaskRunning
		t.action()
	}
}

func (t *task) StopAction() {
	t.state = TaskStopped
	if t.stopAction != nil {
		t.stopAction()
	}
}

func (t *task) GetAlarmTime() time.Time {
	return t.alarmTime
}

func (t *task) GetAction() func() {
	return t.action
}

func (t *task) GetName() string {
	return t.name
}

func (t *task) GetDescription() string {
	return t.description
}

func (t *task) IsTemp() bool {
	return t.isTemp
}

func (t *task) GetState() TaskState {
	return t.state
}

func (t *task) SetState(state TaskState) {
	t.state = state
}

func NewTask(action func(), alarmTime time.Time, name string, description string) (*task, error) {
	return NewTaskWithStopAction(action, nil, alarmTime, name, description, false)
}

func NewTaskWithStopAction(action func(), stopAction func(), alarmTime time.Time, name string, description string, isTemp bool) (*task, error) {
	if action == nil {
		return nil, errors.New("action cannot be nil")
	}
	if alarmTime.Before(time.Now()) {
		return nil, errors.New("alarm time cannot be in the past")
	}
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if description == "" {
		return nil, errors.New("description cannot be empty")
	}
	return &task{
		action:      action,
		stopAction:  stopAction,
		alarmTime:   alarmTime,
		name:        name,
		description: description,
		isTemp:      isTemp,
		state:       TaskPending,
	}, nil
}

package qtodo

import (
	"errors"
	"time"
)

type Task interface {
	DoAction()
	GetAlarmTime() time.Time
	GetAction() func()
	GetName() string
	GetDescription() string
}

type task struct {
	action      func()
	alarmTime   time.Time
	name        string
	description string
}

func (t *task) DoAction() {
	t.action()
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

func NewTask(action func(), alarmTime time.Time, name string, description string) (*task, error) {
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
		alarmTime:   alarmTime,
		name:        name,
		description: description,
	}, nil
}

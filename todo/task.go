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

type MyTask struct {
	Action           func()
	AlarmTime        time.Time
	AlarmName        string
	AlarmDescription string
}

func (task *MyTask) DoAction() {
	task.Action()
}
func (task *MyTask) GetAlarmTime() time.Time { return task.AlarmTime }
func (task *MyTask) GetAction() func()       { return task.Action }
func (task *MyTask) GetName() string         { return task.AlarmName }
func (task *MyTask) GetDescription() string  { return task.AlarmDescription }

func NewTask(action func(), alarmTime time.Time, name string, description string) (Task, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	if time.Now().After(alarmTime) {
		return nil, errors.New("alarm time expired")
	}
	if description == "" {
		return nil, errors.New("description is empty")
	}
	return &MyTask{Action: action,
		AlarmTime:        alarmTime,
		AlarmName:        name,
		AlarmDescription: description,
	}, nil
}

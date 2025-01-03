package qtodo

import (
	"errors"
)

type Database interface {
	GetTaskList() []Task
	GetTask(string) (Task, error)
	SaveTask(Task) error
	DelTask(string) error
}
type InMemoryRepository struct {
	Tasks map[string]*Task
}

func NewDatabase() Database {

	return &InMemoryRepository{Tasks: make(map[string]*Task)}
}

func (d *InMemoryRepository) GetTaskList() []Task {
	result := make([]Task, 0)
	for _, task := range d.Tasks {
		result = append(result, *task)
	}
	return result
}

func (d *InMemoryRepository) GetTask(name string) (Task, error) {
	task, ok := d.Tasks[name]
	if !ok {
		return nil, errors.New("task not found")
	}
	return *task, nil
}

func (d *InMemoryRepository) SaveTask(task Task) error {
	d.Tasks[task.GetName()] = &task
	return nil
}

func (d *InMemoryRepository) DelTask(name string) error {
	delete(d.Tasks, name)
	return nil
}

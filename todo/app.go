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
	GetActiveTaskList() []Task
	GetTask(string) (Task, error)
}

type MyApp struct {
	activeTaskList []Task
	tempTasks      []Task
	tasks          Database
}

func (app *MyApp) StartTask(taskName string) error {
	task, err := app.tasks.GetTask(taskName)
	if err != nil {
		return err
	}
	app.activeTaskList = append(app.activeTaskList, task)
	return nil
}
func (app *MyApp) StopTask(task string) {
	for i, t := range app.activeTaskList {
		if t.GetName() == task {
			app.activeTaskList = append(app.activeTaskList[:i], app.activeTaskList[i+1:]...)
		}
	}
}
func (app *MyApp) GetTaskList() []Task {
	return app.tasks.GetTaskList()
}
func (app *MyApp) GetActiveTaskList() []Task {
	return app.activeTaskList
}
func (app *MyApp) DelTask(task string) error {
	return app.tasks.DelTask(task)
}
func (app *MyApp) GetTask(task string) (Task, error) {
	return app.tasks.GetTask(task)
}
func (app *MyApp) AddTask(taskName string, taskDesc string, t time.Time, f func(), isTemp bool) error {
	task, err := NewTask(f, t, taskName, taskDesc)
	if err != nil {
		return err
	}
	if isTemp {
		app.tempTasks = append(app.tempTasks, task)
	}
	return app.tasks.SaveTask(task)
}

func NewApp(tasks Database) App {
	app := &MyApp{tasks: tasks}
	go app.Run()
	return app
}

func (app *MyApp) Run() {
	for {
		now := time.Now()
		for _, task := range app.activeTaskList {
			if task.GetAlarmTime().Before(now) {
				task.DoAction()
				app.StopTask(task.GetName())
				for i, tempTask := range app.tempTasks {
					if tempTask.GetName() == task.GetName() {
						app.tasks.DelTask(tempTask.GetName())
						app.tempTasks = append(app.tempTasks[:i], app.tempTasks[i+1:]...)
					}
				}
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

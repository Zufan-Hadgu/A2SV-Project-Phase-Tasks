package Domain



type ITaskRepository interface {
	GetTaskByID(taskID string) (Task, error)
	AddTask(task Task) error
	UpdateTask(taskID string, updatedTask Task) error
	DeleteTask(taskID string) error
	GetAllTask() ([]Task, error)
}
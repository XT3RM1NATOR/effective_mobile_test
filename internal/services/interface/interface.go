package infrastructureInterface

import "github.com/XT3RM1NAT0R/time-tracker/internal/domain/entity"

type TaskTrackerRepository interface {
	AddUser(user *entity.User) error
	GetUserById(userId int) (*entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUserById(userId int) error

	AddTask(task entity.Task) error
	GetTasksByUserId(userId, limit, offset int) ([]entity.Task, error)
	GetTaskById(taskId int) (*entity.Task, error)
	UpdateTask(task *entity.Task) error
	DeleteTaskById(taskId int) error
	FinishTask(taskId int) error
}

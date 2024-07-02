package _interface

import (
	model "github.com/XT3RM1NAT0R/time-tracker/internal/delivery/models"
	"github.com/XT3RM1NAT0R/time-tracker/internal/domain/entity"
)

type TaskTrackerService interface {
	AddUser(passportNumber string) (*entity.User, error)
	GetUserById(userId int) (*entity.User, error)
	UpdateUser(userId int, passportNumber string) (*entity.User, error)
	DeleteUserById(userId int) error

	AddTask(userId int, taskName string) (*entity.Task, error)
	GetTasksByUserId(userId, limit, offset int) ([]model.TaskResponse, error)
	UpdateTask(taskId int, taskName string) (*entity.Task, error)
	DeleteTaskById(taskID int) error
	FinishTask(taskId int) error
}

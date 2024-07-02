package services

import (
	"github.com/XT3RM1NAT0R/time-tracker/config"
	model "github.com/XT3RM1NAT0R/time-tracker/internal/delivery/models"
	"github.com/XT3RM1NAT0R/time-tracker/internal/domain/entity"
	_interface "github.com/XT3RM1NAT0R/time-tracker/internal/domain/interface"
	infrastructureInterface "github.com/XT3RM1NAT0R/time-tracker/internal/services/interface"
	"github.com/XT3RM1NAT0R/time-tracker/utils"
	"time"
)

type TaskTrackerServiceImpl struct {
	taskTrackerRepo infrastructureInterface.TaskTrackerRepository
	config          *config.Config
}

func NewTaskTrackerServiceImpl(cfg *config.Config, taskTrackerRepo infrastructureInterface.TaskTrackerRepository) _interface.TaskTrackerService {
	return &TaskTrackerServiceImpl{
		taskTrackerRepo: taskTrackerRepo,
		config:          cfg,
	}
}

func (s *TaskTrackerServiceImpl) AddUser(passportNumber string) (*entity.User, error) {
	name, surname, address, err := utils.FetchRandomUser(s.config.API.EnrichAPI)
	user := &entity.User{
		Name:           name,
		Surname:        surname,
		Address:        address,
		PassportNumber: passportNumber,
	}

	if err = s.taskTrackerRepo.AddUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *TaskTrackerServiceImpl) GetUserById(userId int) (*entity.User, error) {
	user, err := s.taskTrackerRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *TaskTrackerServiceImpl) UpdateUser(userId int, passportNumber string) (*entity.User, error) {
	user, err := s.taskTrackerRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	user.PassportNumber = passportNumber
	err = s.taskTrackerRepo.UpdateUser(*user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *TaskTrackerServiceImpl) DeleteUserById(userId int) error {
	err := s.taskTrackerRepo.DeleteUserById(userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskTrackerServiceImpl) AddTask(userId int, taskName string) (*entity.Task, error) {
	task := &entity.Task{
		UserId:    userId,
		TaskName:  taskName,
		StartTime: time.Now(),
	}

	err := s.taskTrackerRepo.AddTask(*task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskTrackerServiceImpl) GetTasksByUserId(userId, limit, offset int) ([]model.TaskResponse, error) {
	tasks, err := s.taskTrackerRepo.GetTasksByUserId(userId, limit, offset)
	if err != nil {
		return nil, err
	}

	var taskResponses []model.TaskResponse
	for _, task := range tasks {
		taskResponse := model.TaskResponse{
			TaskId:    task.Id,
			UserId:    task.UserId,
			TaskName:  task.TaskName,
			StartTime: task.StartTime,
			EndTime:   task.EndTime,
		}
		taskResponses = append(taskResponses, taskResponse)
	}

	return taskResponses, nil
}

func (s *TaskTrackerServiceImpl) UpdateTask(taskId int, taskName string) (*entity.Task, error) {
	task, err := s.taskTrackerRepo.GetTaskById(taskId)
	if err != nil {
		return nil, err
	}

	task.TaskName = taskName
	err = s.taskTrackerRepo.UpdateTask(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskTrackerServiceImpl) DeleteTaskById(taskId int) error {
	err := s.taskTrackerRepo.DeleteTaskById(taskId)
	if err != nil {
		return err
	}
	return nil
}

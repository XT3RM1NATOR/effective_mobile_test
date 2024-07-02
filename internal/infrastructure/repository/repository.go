package repository

import (
	"github.com/XT3RM1NAT0R/time-tracker/config"
	"github.com/XT3RM1NAT0R/time-tracker/internal/domain/entity"
	infrastructureInterface "github.com/XT3RM1NAT0R/time-tracker/internal/services/interface"

	"github.com/jmoiron/sqlx"
)

type TaskTrackerRepositoryImpl struct {
	database *sqlx.DB
	config   *config.Config
}

func NewTaskTrackerRepositoryImpl(cfg *config.Config, db *sqlx.DB) infrastructureInterface.TaskTrackerRepository {
	return &TaskTrackerRepositoryImpl{
		database: db,
		config:   cfg,
	}
}

func (r *TaskTrackerRepositoryImpl) AddUser(user *entity.User) error {
	query := `
		INSERT INTO users (passport_number, name, surname, patronymic, address)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING Id
	`

	return r.database.QueryRowx(query, user.PassportNumber, user.Name, user.Surname, user.Patronymic, user.Address).Scan(&user.Id)
}

func (r *TaskTrackerRepositoryImpl) GetUserById(userId int) (*entity.User, error) {
	var user entity.User
	query := `SELECT * FROM users WHERE Id = $1`

	if err := r.database.Get(&user, query, userId); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *TaskTrackerRepositoryImpl) UpdateUser(user entity.User) error {
	query := `
		UPDATE users
		SET passport_number = $1, name = $2, surname = $3, patronymic = $4, address = $5
		WHERE Id = $6
	`

	_, err := r.database.Exec(query, user.PassportNumber, user.Name, user.Surname, user.Patronymic, user.Address, user.Id)
	return err
}

func (r *TaskTrackerRepositoryImpl) DeleteUserById(userId int) error {
	query := `DELETE FROM users WHERE Id = $1`

	_, err := r.database.Exec(query, userId)
	return err
}

func (r *TaskTrackerRepositoryImpl) AddTask(task entity.Task) error {
	query := `
		INSERT INTO tasks (user_Id, task_name, start_time, end_time)
		VALUES ($1, $2, $3, $4)
		RETURNING Id
	`

	return r.database.QueryRowx(query, task.UserId, task.TaskName, task.StartTime, task.EndTime).Scan(&task.Id)
}

func (r *TaskTrackerRepositoryImpl) GetTasksByUserId(userId, limit, offset int) ([]entity.Task, error) {
	var tasks []entity.Task
	query := `
		SELECT * FROM tasks 
		WHERE user_id = $1 
		ORDER BY 
			CASE 
				WHEN end_time IS NULL THEN 0 
				ELSE 1 
			END, 
			(end_time - start_time) DESC 
		LIMIT $2 OFFSET $3`

	err := r.database.Select(&tasks, query, userId, limit, offset)
	return tasks, err
}

func (r *TaskTrackerRepositoryImpl) GetTaskById(taskId int) (*entity.Task, error) {
	var task entity.Task
	query := `SELECT * FROM tasks WHERE Id = $1`

	if err := r.database.Get(&task, query, taskId); err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskTrackerRepositoryImpl) UpdateTask(task *entity.Task) error {
	query := `
		UPDATE tasks
		SET user_Id = $1, task_name = $2, start_time = $3, end_time = $4
		WHERE Id = $5
	`

	_, err := r.database.Exec(query, task.UserId, task.TaskName, task.StartTime, task.EndTime, task.Id)
	return err
}

func (r *TaskTrackerRepositoryImpl) DeleteTaskById(taskId int) error {
	query := `DELETE FROM tasks WHERE Id = $1`

	_, err := r.database.Exec(query, taskId)
	return err
}

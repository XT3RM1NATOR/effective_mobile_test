package controllers

import (
	"github.com/XT3RM1NAT0R/time-tracker/config"
	model "github.com/XT3RM1NAT0R/time-tracker/internal/delivery/models"
	_interface "github.com/XT3RM1NAT0R/time-tracker/internal/domain/interface"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TaskTrackerController struct {
	taskTrackerService _interface.TaskTrackerService
	config             *config.Config
}

func NewTaskTrackerController(cfg *config.Config, taskTrackerService _interface.TaskTrackerService) *TaskTrackerController {
	return &TaskTrackerController{
		taskTrackerService: taskTrackerService,
		config:             cfg,
	}
}

// AddUserHandler adds a new user.
// @Summary Add a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body model.AddUserRequest true "User details"
// @Success 200 {object} model.SuccessResponse "User added successfully"
// @Failure 400 {object} model.ErrorResponse "Invalid request payload"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /users [post]
func (c *TaskTrackerController) AddUserHandler(ctx echo.Context) error {
	var req model.AddUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid request payload"})
	}

	_, err := c.taskTrackerService.AddUser(req.PassportNumber)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, model.SuccessResponse{Message: "user added successfully"})
}

// GetUserByIdHandler retrieves a user by Id.
// @Summary Get user by Id
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} model.UserResponse "User details"
// @Failure 400 {object} model.ErrorResponse "Invalid user ID"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /users/{userId} [get]
func (c *TaskTrackerController) GetUserByIdHandler(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid user ID"})
	}

	user, err := c.taskTrackerService.GetUserById(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, model.UserResponse{
		Id:             user.Id,
		PassportNumber: user.PassportNumber,
		Name:           user.Name,
		Surname:        user.Surname,
		Patronymic:     user.Patronymic,
		Address:        user.Address,
	})
}

// UpdateUserHandler updates a user's details.
// @Summary Update user details
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path int true "User Id"
// @Param request body model.UpdateUserRequest true "Updated user details"
// @Success 200 {object} model.SuccessResponse "User updated successfully"
// @Failure 400 {object} model.ErrorResponse "Invalid request payload or user ID"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /users/{userId} [put]
func (c *TaskTrackerController) UpdateUserHandler(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid user ID"})
	}

	var req model.UpdateUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid request payload"})
	}

	if _, err = c.taskTrackerService.UpdateUser(userID, req.PassportNumber); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, model.SuccessResponse{Message: "user updated successfully"})
}

// DeleteUserByIdHandler deletes a user by Id.
// @Summary Delete user by Id
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} model.SuccessResponse "User deleted successfully"
// @Failure 400 {object} model.ErrorResponse "Invalid user ID"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /users/{userId} [delete]
func (c *TaskTrackerController) DeleteUserByIdHandler(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid user ID"})
	}

	if err = c.taskTrackerService.DeleteUserById(userId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, model.SuccessResponse{Message: "user deleted successfully"})
}

// AddTaskHandler adds a new task.
// @Summary Add a new task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param request body model.AddTaskRequest true "Task details"
// @Success 200 {object} model.SuccessResponse "Task created successfully"
// @Failure 400 {object} model.ErrorResponse "Invalid request payload"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /tasks [post]
func (c *TaskTrackerController) AddTaskHandler(ctx echo.Context) error {
	var req model.AddTaskRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid request payload"})
	}

	if _, err := c.taskTrackerService.AddTask(req.UserId, req.TaskName); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, model.SuccessResponse{Message: "task created successfully"})
}

// GetTasksByUserIdHandler retrieves tasks for a user with pagination.
// @Summary Get tasks by user ID with pagination
// @Tags Tasks
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} model.TaskResponse "List of tasks"
// @Failure 400 {object} model.ErrorResponse "Invalid user ID or pagination parameters"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /tasks/{userId} [get]
func (c *TaskTrackerController) GetTasksByUserIdHandler(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid user Id"})
	}

	//пагинация
	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	tasks, err := c.taskTrackerService.GetTasksByUserId(userId, limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, tasks)
}

// UpdateTaskHandler updates a task's details.
// @Summary Update task details
// @Tags Tasks
// @Accept json
// @Produce json
// @Param request body model.UpdateTaskRequest true "Updated task details"
// @Success 200 {object} model.TaskResponse "Task updated successfully"
// @Failure 400 {object} model.ErrorResponse "Invalid request payload"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /tasks [put]
func (c *TaskTrackerController) UpdateTaskHandler(ctx echo.Context) error {
	var req model.UpdateTaskRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid request payload"})
	}

	task, err := c.taskTrackerService.UpdateTask(req.TaskId, req.Name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, task)
}

// DeleteTaskByIdHandler deletes a task by Id.
// @Summary Delete task by Id
// @Tags Tasks
// @Accept json
// @Produce json
// @Param taskId path int true "Task ID"
// @Success 200 {object} model.SuccessResponse "Task deleted successfully"
// @Failure 400 {object} model.ErrorResponse "Invalid task ID"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /tasks/{taskId} [delete]
func (c *TaskTrackerController) DeleteTaskByIdHandler(ctx echo.Context) error {
	taskId, err := strconv.Atoi(ctx.Param("taskId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid task ID"})
	}

	if err = c.taskTrackerService.DeleteTaskById(taskId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}

	return ctx.JSON(http.StatusOK, model.SuccessResponse{Message: "task deleted successfully"})
}

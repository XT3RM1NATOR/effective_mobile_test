package delivery

import (
	"github.com/XT3RM1NAT0R/time-tracker/config"
	"github.com/XT3RM1NAT0R/time-tracker/internal/delivery/controllers"
	"github.com/XT3RM1NAT0R/time-tracker/internal/infrastructure/repository"
	"github.com/XT3RM1NAT0R/time-tracker/internal/services"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, cfg *config.Config, db *sqlx.DB) {
	ttr := repository.NewTaskTrackerRepositoryImpl(cfg, db)
	tts := services.NewTaskTrackerServiceImpl(cfg, ttr)
	ic := controllers.NewTaskTrackerController(cfg, tts)

	e.POST("/users", ic.AddUserHandler)
	e.GET("/users/:userId", ic.GetUserByIdHandler)
	e.PUT("/users/:userId", ic.UpdateUserHandler)
	e.DELETE("/users/:userId", ic.DeleteUserByIdHandler)

	e.POST("/tasks", ic.AddTaskHandler)
	e.GET("/tasks/:userId", ic.GetTasksByUserIdHandler)
	e.PUT("/tasks", ic.UpdateTaskHandler)
	e.DELETE("/tasks/:taskId", ic.DeleteTaskByIdHandler)
}

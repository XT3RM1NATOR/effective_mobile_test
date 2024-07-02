package model

import (
	"time"
)

// Requests

type AddUserRequest struct {
	PassportNumber string `json:"passportNumber"`
}

type UpdateUserRequest struct {
	UserId         int    `json:"userId"`
	PassportNumber string `json:"passportNumber"`
}

type UpdateTaskRequest struct {
	TaskId     int    `json:"taskId"`
	Name       string `json:"name"`
	IsFinished bool   `json:"isFinished"`
}

type AddTaskRequest struct {
	UserId   int    `json:"userId"`
	TaskName string `json:"taskName"`
}

type GetTasksRequest struct {
	UserId int `json:"userId"`
}

// Responses

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type TaskResponse struct {
	TaskId    int       `json:"taskId"`
	UserId    int       `json:"userId"`
	TaskName  string    `json:"taskName"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type DeleteTaskResponse struct {
	TaskId int `json:"taskId"`
}

type UserResponse struct {
	Id             int    `json:"id"`
	PassportNumber string `json:"passportNumber"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}

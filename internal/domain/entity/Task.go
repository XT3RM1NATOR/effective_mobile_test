package entity

import "time"

type Task struct {
	Id        int       `db:"id"`
	UserId    int       `db:"user_id"`
	TaskName  string    `db:"task_name"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
}

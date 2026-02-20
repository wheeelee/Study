package sql_test

import "time"

type TaskModel struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

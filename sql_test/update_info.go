package sql_test

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateInfo(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	UPDATE tasks
	SET completed = TRUE
	WHERE id = 2 OR id = 4;
	`
	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
func UpdateTask(
	ctx context.Context,
	conn *pgx.Conn,
	task TaskModel,
) error {
	sqlQuery := `
	UPDATE tasks
	SET title = $1,description = $2,completed = $3,created_at = $4,completed_at = $5
	WHERE id = $6;
	`
	_, err := conn.Exec(
		ctx,
		sqlQuery,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.CompletedAt,
		task.ID,
	)
	return err
}

package main

import (
	"context"
	"fmt"
	"study/simple_connection"
	"study/sql_test"
	"time"

	"github.com/k0kubun/pp"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := sql_test.Create_table(ctx, conn); err != nil {
		panic(err)
	}

	// if err := sql_test.Insert_Row(
	// 	ctx,
	// 	conn,
	// 	"Полдник",
	// 	"Надо поесть",
	// 	false,
	// 	time.Now(),
	// ); err != nil {
	// 	panic(err)
	// }

	// if err := sql_test.DeleteRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// if err := sql_test.UpdateInfo(ctx, conn); err != nil {
	// 	panic(err)
	// }

	tasks, err := sql_test.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}
	for _, task := range tasks {
		if task.ID == 3 {
			task.Title = "Кошка"
			task.Description = "Люблю своих котов"
			task.Completed = true
			now := time.Now()
			task.CompletedAt = &now

			if err := sql_test.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}
			break
		}
	}
	pp.Println(tasks)
	fmt.Println("succed!")
}

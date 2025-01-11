package db

import (
	"context"
	"todo_list_telegram/model"
	"todo_list_telegram/postgresql"
)

type TodoDb struct {
	pg postgresql.PgClient
}

func NewTodoDb(pg postgresql.PgClient) *TodoDb {
	return &TodoDb{
		pg: pg,
	}
}

func (d *TodoDb) AddTodo(c context.Context, todo model.Todo) error {
	err := d.pg.QueryRow(c, `INSERT INTO todos (title, done) VALUES ($1, $2) RETURNING uuid`, todo.Title, todo.Done).Scan(&todo.UUID)
	if err != nil {
		return err
	}
	return nil
}

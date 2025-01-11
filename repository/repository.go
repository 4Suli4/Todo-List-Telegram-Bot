package repository

import (
	"context"
	"todo_list_telegram/model"
)

type Repository interface {
	AddTodo(c context.Context, todo model.Todo) error
}

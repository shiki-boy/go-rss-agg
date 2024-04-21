package main

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shiki-boy/go-rss-agg/internal/database"
)

type Todo struct {
	ID        int32            `json:"id"`
	Title     pgtype.Text      `json:"title"`
	Done      pgtype.Bool      `json:"done"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func formatTodoDto(dbTodo database.Todo) Todo {
	return Todo{
		ID:        dbTodo.ID,
		Title:     dbTodo.Title,
		Done:      dbTodo.Done,
		CreatedAt: dbTodo.CreatedAt,
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shiki-boy/go-rss-agg/internal/database"
)

func (apiCfg *apiConfig) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Title pgtype.Text `json:"title"`
		Done  pgtype.Bool `json:"done"`
	}
	params := parameters{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error in parsing data: %s", err))
		return
	}

	newTodo, err := apiCfg.DB.CreateTodo(r.Context(), database.CreateTodoParams{
		Title: params.Title,
		Done:  params.Done,
	})

	if err != nil {
		respondWithErr(w, 400, fmt.Sprintf("Error in creating todo: %s", err))
		return
	}

	respondWithJson(w, 201, formatTodoDto(newTodo))
}

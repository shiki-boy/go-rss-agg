package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/shiki-boy/go-rss-agg/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("No PORT provided! Exiting...")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))

	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}
	log.Println("Successfully connected to DB")

	defer conn.Close(context.Background())

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.CleanPath) // ? users////1 will both be treated as: /users/1

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Recoverer) // ! recovers from panics, logs the panic and returns 500

	log.Println("running on PORT:" + port)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		type res struct {
			Msg string `json:"message"`
		}

		respondWithJson(w, 200, res{
			"Hello world",
		})
	})

	r.Get("/err", func(w http.ResponseWriter, r *http.Request) {
		respondWithErr(w, 400, "Sample error")
	})

	r.Post("/users", apiCfg.handleCreateTodo)

	err = http.ListenAndServe(":"+port, r)

	if err != nil {
		log.Fatal("Server crashed: ", err)
	}
}

# RSS agg.

1. go mod init github.com/shiki-boy/go-rss-agg
2. go get github.com/joho/godotenv
3. go mod tidy
4. go mod vendor # get local copy of the libraries, run everytime when you add new 
5. go get github.com/go-chi/chi
6. go get github.com/go-chi/cors
7. added server restart with nodemon
8. set -a && source .env && set +a # for setting env var for goose
9. goose create <migration_name> sql
10. goose up
11. goose down
12. sqlc generate
13. created sqlc.yml
14. internal/database is code generated by sqlc
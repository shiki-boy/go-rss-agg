# RSS agg.

go mod init github.com/shiki-boy/go-rss-agg
go get github.com/joho/godotenv
go mod tidy
go mod vendor # get local copy of the libraries, run everytime when you add new 

go get github.com/go-chi/chi
go get github.com/go-chi/cors
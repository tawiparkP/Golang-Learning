module github.com/tawiparkP/railAPI-Gin

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/mattn/go-sqlite3 v1.14.3
)

require github.com/tawiparkP/dbutils v0.0.0

replace github.com/tawiparkP/dbutils => ../dbutils

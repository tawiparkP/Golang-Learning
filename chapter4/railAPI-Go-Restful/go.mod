module github.com/tawiparkP/railAPI

go 1.15

require (
	github.com/emicklei/go-restful v2.14.2+incompatible
	github.com/mattn/go-sqlite3 v1.14.3
)

require github.com/tawiparkP/dbutils v0.0.0
replace github.com/tawiparkP/dbutils => ../dbutils

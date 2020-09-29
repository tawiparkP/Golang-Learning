package main

import (
	"log"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type Book struct {
    id int
    name string
    author string
}

func dbOperaions(db *sql.DB){
    statement, _ := db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?,?,?)")
    statement.Exec("Dacula","Bram Stoker",14043547)
    log.Println("Ïnserted the book into database")

    rows, _ := db.Query("SELECT id, name, author FROM books")
    var tempBook Book
    for rows.Next(){
        rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
        log.Printf("ID:%d, Book:%s, Author:%s\n ", tempBook.id, tempBook.name, tempBook.author)
    }
}

func main(){
    db, err := sql.Open("sqlite3", "./books.db")
    if err != nil {
        log.Fatal(err)
    }

    statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
    if err != nil {
        log.Println("Error in creating table : books")
    }else{
        log.Println("Successfully created table : books")
    }

    statement.Exec()
    dbOperaions(db)
}


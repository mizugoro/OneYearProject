package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func execDB(db *sql.DB, q string) {
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err)
	}
}

func main() {

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	q := `
        CREATE TABLE memo (
          id INTEGER PRIMARY KEY AUTOINCREMENT,
          body VARCHAR(255) NOT NULL,
          created_at TIMESTAMP DEFAULT (DATETIME('now','localtime'))
        )
    `
	execDB(db, q)

	q = `
        INSERT INTO memo(body)
        VALUES ('body1'), ('body2')
    `
	execDB(db, q)

	db.Close()
}

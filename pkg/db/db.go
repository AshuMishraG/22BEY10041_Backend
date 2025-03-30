package db

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func Connect() *sql.DB {
    connStr := "user=postgres dbname=file_sharing_db sslmode=disable" // Update with your credentials
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    return db
}
// Chapter 14: Database Access: SQL (database/sql)
package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/lib/pq"
)

func main() {
    dsn := "host=localhost port=5432 user=pguser password=secret dbname=mydb sslmode=disable"
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(30 * time.Minute)
    db.Ping()

    var name string
    err = db.QueryRow("SELECT name FROM users WHERE id=$1", 1).Scan(&name)
    if err != nil {
        fmt.Println("Query error:", err)
    } else {
        fmt.Println("User name:", name)
    }
}

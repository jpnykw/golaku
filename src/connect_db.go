package main

import (
  "fmt"
  "os"
  "database/sql"
  "github.com/joho/godotenv"
  _ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
  err := godotenv.Load()
  if err != nil {
    fmt.Println(err.Error())
  }

  user := os.Getenv("DB_USER")
  pass := os.Getenv("DB_PASS")
  host := os.Getenv("DB_HOST")
  port := os.Getenv("DB_PORT")
  name := os.Getenv("DB_NAME")
  query := "?charset=utf8mb4"

  dbconf := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + query
  db, err := sql.Open("mysql", dbconf)
  if err != nil {
    fmt.Println(err.Error())
  }

  return db
}

func main() {
  db := connect()
  defer db.Close()

  if db.Ping() != nil {
    fmt.Println("Failed to connect to the DB.")
    return
  }
  fmt.Println("Successfully connected to the DB.")

  // testing simple query
  res, _ := db.Query("SHOW TABLES")
  var tables string
  for res.Next() {
    res.Scan(&tables)
    fmt.Println(tables)
  }
}


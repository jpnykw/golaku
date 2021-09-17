package main

import (
  "fmt"
  "os"

  "net/http"
  "github.com/labstack/echo"

  "database/sql"
  "github.com/joho/godotenv"
  _ "github.com/go-sql-driver/mysql"
)

type User struct {
  Name  string `json:"name"`
  Age   string `json:"age"`
  Email string `json:"email"`
}

func main() {
  // Listen
  e := echo.New()
  e.GET("/reverse/:text", reverse)
  e.GET("/users/get/:name", get_user)
  e.POST("/users/create", create_user)
  e.Logger.Fatal(e.Start(":5678"))
}

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

func reverse_string(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func reverse(ctx echo.Context) error {
  text := ctx.Param("text")
  return ctx.String(http.StatusOK, reverse_string(text))
}

func get_user(ctx echo.Context) error {
  // Connect
  db := connect()
  if db.Ping() != nil {
    return ctx.String(http.StatusInternalServerError, "Failed to connect to the DB.")
  }

  // GET
  name := ctx.Param("name")
  res, _ := db.Query("select * from users where name='" + name + "'")

  result := "data not found"
  for res.Next() {
    var id int
    var name string
    var age int
    var email string
    res.Scan(&id, &name, &age, &email)
    result = fmt.Sprintf("id: %d, name: %s, age: %d, email: %s", id, name, age, email)
  }

  db.Close()
  return ctx.String(http.StatusOK, result)
}


func create_user(ctx echo.Context) error {
  // Connect
  db := connect()
  if db.Ping() != nil {
    return ctx.String(http.StatusInternalServerError, "Failed to connect to the DB.")
  }

  // POST
  user := new(User)
  if err := ctx.Bind(user); err != nil {
    return err
  }
  
  name := user.Name
  age := user.Age
  email := user.Email

  ins, err := db.Prepare("insert into users (name, age, email) values (?,?,?)")
  if err != nil {
    return err
  }

  ins.Exec(name, age, email)
  return ctx.JSON(http.StatusOK, user)
}


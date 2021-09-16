package main

import (
  "net/http"
  "github.com/labstack/echo"
)

type User struct {
  Name  string `json:"name"`
  Age   string `json:"age"`
  Email string `json:"email"`
}

func main() {
  e := echo.New()
  e.GET("/reverse/:text", reverse)
  e.POST("/signup", signup)
  e.Logger.Fatal(e.Start(":5678"))
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

func signup(ctx echo.Context) error {
  user := new(User)
  if err := ctx.Bind(user); err != nil {
    return err
  }
  return ctx.JSON(http.StatusOK, user)
}

package models

import (
  "fmt"
  "log"
  "net/http"

  "github.com/labstack/echo"

  "github.com/Tyler-Scheerens/cdm/server/datasources"
)

type User struct {
  Id       int64
  Username string
  Password string
  Salt     string
  Email    string
}

type DBUser struct {
  Id       int64  `db:"id"       json:"id"`
  Username string `db:"username" json:"username"`
  Password string `db:"password" json:"-"`
  Salt     string `db:"salt"     json:"-"`
  Email    string `db:"email"    json:"email"`
}

func checkError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

// GET /api/user/:id
func GetUserById(c echo.Context) error {
  id := c.Param("id")
  db := datasources.GetDb()

  fmt.Println(id)
  dbUser := DBUser{}
  err := db.Get(&dbUser, "SELECT * FROM user where id = ?", id)
  checkError(err)

  return c.JSON(http.StatusOK, dbUser)
}

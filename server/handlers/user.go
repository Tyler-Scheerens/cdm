package handlers

import (
  "net/http"

  "github.com/labstack/echo"

  "github.com/Tyler-Scheerens/cdm/server/datasources"
  "github.com/Tyler-Scheerens/cdm/server/models"
)

// GET /api/user/:id
func GetUserById(c echo.Context) error {
  id := c.Param("id")
  db := datasources.GetDb()

  login := models.Login{}
  err := db.Get(&login, "SELECT * FROM login WHERE user_id = $1", id)
  if err != nil {
    return c.String(http.StatusBadRequest, err.Error())
  }

  return c.JSON(http.StatusOK, login)
}
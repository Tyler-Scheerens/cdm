package handlers

import (
  "github.com/labstack/echo"

  "github.com/Tyler-Scheerens/cdm/server/models"
)

func Register(e *echo.Echo) {
  e.GET("/api/user/:id", models.GetUserById)
}

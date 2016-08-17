package handlers

import (
  "github.com/labstack/echo"

  "github.com/Tyler-Scheerens/cdm/server/handlers"
)

func Register(e *echo.Echo) {
  e.GET("/api/user/:id", models.GetUserById)
}

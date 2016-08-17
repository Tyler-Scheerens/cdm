package handlers

import (
  "github.com/labstack/echo"

  "../models"
)

func Register(e *echo.Echo) {
  e.GET("/api/user/:id", models.GetUserById)
}

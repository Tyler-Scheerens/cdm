package handlers

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"

  "github.com/Tyler-Scheerens/cdm/server/conf"
  "github.com/Tyler-Scheerens/cdm/server/models"
)

var serverConfig = conf.Server{}

func Authenticate(username, password string) bool {
  if username == serverConfig.Username && password == serverConfig.Password {
    return true
  }
  return false
}

func Register(c *conf.Config, e *echo.Echo) {
  serverConfig = c.Server
  
  e.Use(middleware.BasicAuth(Authenticate))

  e.GET("/api/user/:id", models.GetUserById)
}

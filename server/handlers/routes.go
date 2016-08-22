package handlers

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"

  "github.com/Tyler-Scheerens/cdm/server/conf"
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
  
  e.File("/", serverConfig.UIRoot + "src/index.html")
  e.File("/index.js", serverConfig.UIRoot + "dist/client/index.js")
  e.File("/home", serverConfig.UIRoot + "src/index.html");
  e.File("/home/*", serverConfig.UIRoot + "src/index.html");
  e.File("/login", serverConfig.UIRoot + "src/index.html");
  e.File("/login/*", serverConfig.UIRoot + "src/index.html");

  e.Static("/", serverConfig.UIRoot + "src")
  e.Static("/node_modules", serverConfig.UIRoot + "node_modules")

  // alert
  e.GET("/api/alert/:id", GetAlertById)

  // user
  e.GET("/api/user/:id", GetUserById)
}

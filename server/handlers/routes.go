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
  
  e.File("/", serverConfig.UIRoot + "dist/index.html")

  e.Static("/", serverConfig.UIRoot + "dist")
  e.Static("/node_modules", serverConfig.UIRoot + "node_modules")

  // user
  e.PUT("/api/user", CreateUser)
  e.GET("/api/user/:id", GetUserById)
  e.GET("/api/user/login", UserLogin)
  e.GET("/api/user/logout", UserLogout)

  // alert
  e.GET("/api/alert", GetAlert)
  e.PUT("/api/alert", CreateAlert)
  e.GET("/api/alert/:id", GetAlertById)
  e.DELETE("/api/alert/:id", RemoveAlertById)

  // elasticsearch
  e.GET("/api/elasticsearch", GetElasticsearch)
  e.PUT("/api/elasticsearch", CreateElasticsearch)
  e.GET("/api/elasticsearch/:id", GetElasticsearchById)
}

package handlers

import (
  "errors"
  "net/http"

  "github.com/labstack/echo"

  "github.com/Tyler-Scheerens/cdm/server/datasources"
  "github.com/Tyler-Scheerens/cdm/server/models"
)

// PUT /api/user
func CreateUser(c echo.Context) error {
  // TODO encrypt/salt password
  login := models.Login{}
  err := datasources.GetDb().Get(&login,
    "INSERT INTO login (username, password, email) VALUES ($1, $2, $3)",
    c.Param("username"),
    c.Param("password"),
    c.Param("email"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, err.Error())
  }

  return c.JSON(http.StatusOK, login)
}

// GET /api/user/:id
func GetUserById(c echo.Context) error {
  login := models.Login{}
  err := datasources.GetDb().Get(&login, "SELECT * FROM login WHERE user_id = $1", c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, err.Error())
  }

  return c.JSON(http.StatusOK, login)
}

// GET /api/user/login
func UserLogin(c echo.Context) error {
  login := models.Login{}
  err := datasources.GetDb().Get(&login,
    "SELECT * FROM login WHERE user_id = $1",
    c.Param("username"),
    c.Param("password"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, err.Error())
  }

  // TODO encrpt/salt password
  if (login.Password != c.Param("password")) {
    return c.JSON(http.StatusBadRequest, errors.New("Invalid username or password"))
  }

  // TODO create session
  return c.JSON(http.StatusOK, login)
}


// GET /api/user/logout
func UserLogout(c echo.Context) error {
  // TODO
  return c.JSON(http.StatusOK, "")
}
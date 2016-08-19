package models

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/Tyler-Scheerens/cdm/server/datasources"
)

type Login struct {
	UserId   int64  `db:"user_id"  json:"user_id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"-"`
	Salt     string `db:"salt"     json:"-"`
	Email    string `db:"email"    json:"email"`
}

// GET /api/user/:id
func GetUserById(c echo.Context) error {
	id := c.Param("id")
	db := datasources.GetDb()

	login := Login{}
	err := db.Get(&login, "SELECT * FROM login WHERE user_id = $1", id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, login)
}

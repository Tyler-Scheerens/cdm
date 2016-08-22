package handlers

import (
  "net/http"

  "github.com/labstack/echo"

  "github.com/Tyler-Scheerens/cdm/server/datasources"
  "github.com/Tyler-Scheerens/cdm/server/models"
)

// GET /api/alert/:id
func GetAlertById(c echo.Context) error {
  id := c.Param("id")
  db := datasources.GetDb()

  alert := models.Alert{}
  err := db.Get(&alert, "SELECT * FROM alert WHERE alert_id = $1", id)
  if err != nil {
    return c.String(http.StatusBadRequest, err.Error())
  }

  return c.JSON(http.StatusOK, alert)
}
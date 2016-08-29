package handlers

import (
  "fmt"
  "net/http"

  "github.com/labstack/echo"

  "github.com/Tyler-Scheerens/cdm/server/datasources"
  "github.com/Tyler-Scheerens/cdm/server/models"
)

// GET /api/alert
func GetAlert(c echo.Context) error {
  db := datasources.GetDb()

  alerts := []models.Alert{}
  err := db.Select(
    &alerts,
    "SELECT * FROM alert")
  if err != nil {
    return c.JSON(http.StatusBadRequest, err.Error())
  }

  for _, alert := range alerts {
    // TODO
    if alert.RuleType == "any" {
      aa := models.AnyAlert{}
      alert.AlertOptions.Unmarshal(&aa)
      fmt.Println(aa)
    } else if alert.RuleType == "frequency" {
      fa := models.FrequencyAlert{}
      alert.AlertOptions.Unmarshal(&fa)
      fmt.Println(fa)
    }
  }

  return c.JSON(http.StatusOK, alerts)
}

// PUT /api/alert
func CreateAlert(c echo.Context) error {
  db := datasources.GetDb()

  // TODO parse alert_options
  alert := models.Alert{}
  err := db.Get(
    &alert,
    "INSERT (elasticsearch_id, rule_name, rule_type, alert_type) INTO alert VALUES ($1, $2, $3, $4)",
    c.Param("elasticsearch_id"),
    c.Param("rule_name"),
    c.Param("rule_type"),
    c.Param("alert_type"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, err.Error())
  }

  return c.JSON(http.StatusOK, alert)
}

// GET /api/alert/:id
func GetAlertById(c echo.Context) error {
  db := datasources.GetDb()

  alert := models.Alert{}
  err := db.Get(
    &alert,
    "SELECT * FROM alert WHERE alert_id = $1",
    c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, err.Error())
  }

  return c.JSON(http.StatusOK, alert)
}

// DELETE /api/alert/:id
func RemoveAlertById(c echo.Context) error {
  db := datasources.GetDb()

  alert := models.Alert{}
  err := db.Get(
    &alert,
    "DELETE * FROM alert WHERE alert_id = $1",
    c.Param("id"))
  if err == nil {
    return c.JSON(http.StatusBadRequest, err.Error())
  }

  return c.JSON(http.StatusOK, alert)
}
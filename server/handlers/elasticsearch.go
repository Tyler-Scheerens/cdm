package handlers

import (
  "net/http"

  "github.com/labstack/echo"

  "github.com/Tyler-Scheerens/cdm/server/datasources"
  "github.com/Tyler-Scheerens/cdm/server/models"
)

// GET /api/elasticsearch
func GetElasticsearch(c echo.Context) error {
  db := datasources.GetDb()

  ess := []models.Elasticsearch{}
  err := db.Select(
    &ess,
    "SELECT * FROM elasticsearch")
  if err != nil {
    return c.JSON(http.StatusBadRequest, err.Error())
  }

  return c.JSON(http.StatusOK, ess)
}

// PUT /api/elasticsearch
func CreateElasticsearch(c echo.Context) error {
  db := datasources.GetDb()

  // TODO parse alert_options
  es := models.Elasticsearch{}
  err := db.Get(
    &es,
    "INSERT (ip, port, alert_index) INTO elasticsearch VALUES ($1, $2, $3)",
    c.Param("ip"),
    c.Param("port"),
    c.Param("alert_index"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, err.Error())
  }

  return c.JSON(http.StatusOK, es)
}

// GET /api/elasticsearch/:id
func GetElasticsearchById(c echo.Context) error {
  db := datasources.GetDb()

  es := models.Elasticsearch{}
  err := db.Get(
    &es,
    "SELECT * FROM es WHERE elasticsearch_id = $1",
    c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, err.Error())
  }

  return c.JSON(http.StatusOK, es)
}

// DELETE /api/elasticsearch/:id
func RemoveElasticsearchById(c echo.Context) error {
  db := datasources.GetDb()

  es := models.Alert{}
  err := db.Get(
    &es,
    "DELETE * FROM elasticsearch WHERE elasticsearch_id = $1",
    c.Param("id"))
  if err == nil {
    return c.JSON(http.StatusBadRequest, err.Error())
  }

  return c.JSON(http.StatusOK, es)
}
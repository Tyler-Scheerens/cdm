package datasources

import (
  "fmt"
  "log"

  _ "github.com/lib/pq"
  "github.com/jmoiron/sqlx"

  "github.com/Tyler-Scheerens/cdm/server/conf"
)

type DatasourceManager struct {
  Name string
  Bind string
  DB *sqlx.DB
}

var manager = DatasourceManager{}

func checkError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func AddDatasource(config *conf.Config) {
  datasourceBind := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable",
    config.Datasource.Host,
    config.Datasource.Port,
    config.Datasource.Database,
  )

  manager.Name = config.Datasource.Name
  manager.Bind = datasourceBind

  db, err := sqlx.Connect(manager.Name, manager.Bind)
  manager.DB = db
  checkError(err)
}

func GetDb() *sqlx.DB {
  return manager.DB
}

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
}

var manager = DatasourceManager{}

func checkError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func GetDb() *sqlx.DB {
  db, err := sqlx.Connect(manager.Name, manager.Bind)
  checkError(err)
  return db
}

func AddDatasource(config *conf.Config) {
  datasourceBind := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable",
    config.Datasource.Host,
    config.Datasource.Port,
    config.Datasource.Database,
  )

  manager.Name = config.Datasource.Name
  manager.Bind = datasourceBind
}

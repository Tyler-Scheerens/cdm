package datasources

import (
  "fmt"
  "log"

  _ "github.com/go-sql-driver/mysql"
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
  datasourceBind := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
    config.Datasource.Username,
    config.Datasource.Password,
    config.Datasource.Host,
    config.Datasource.Port,
    config.Datasource.Database,
  )

  manager.Name = config.Datasource.Name
  manager.Bind = datasourceBind
}

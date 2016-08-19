package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"

	"github.com/Tyler-Scheerens/cdm/server/conf"
	"github.com/Tyler-Scheerens/cdm/server/datasources"
	"github.com/Tyler-Scheerens/cdm/server/handlers"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	config, err := conf.NewConfig("config.yml").Load()
	checkError(err)

	datasources.AddDatasource(config)

	e := echo.New()
	handlers.Register(config, e)

	serverBind := fmt.Sprintf("%s:%s",
		config.Server.IP,
		config.Server.Port,
	)

	e.Run(fasthttp.New(serverBind))
}

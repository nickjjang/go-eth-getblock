package main

import (
	"ethcache/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.Build(e)
	e.Logger.Fatal(e.Start(":7000"))
}

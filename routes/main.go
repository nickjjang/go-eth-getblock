package routes

import (
	"ethcache/controllers"

	"github.com/labstack/echo/v4"
)

func Build(e *echo.Echo) {
	e.GET("/", controllers.Home)

	e.GET("/block/:block", controllers.Block)
	e.GET("/block/:block/txs/:txs", controllers.Transaction)
}

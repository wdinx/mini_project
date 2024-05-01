package main

import (
	"github.com/labstack/echo/v4"
	"mini_project/config"
	"mini_project/util"
)

func main() {
	util.LoadEnv()
	cfg := config.InitConfigMySQL()
	e := echo.New()

	db := config.InitDB(cfg)

	e.Start(":3000")
}

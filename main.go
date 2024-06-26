package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"mini_project/config"
	"mini_project/route"
	"mini_project/util"
)

func main() {
	util.LoadEnv()
	cfg := config.Get()
	e := echo.New()
	validate := validator.New()

	db := config.InitDB(cfg.Database)

	route.InitRoute(db, e, validate, cfg)

	e.Start(":3000")
}

package main

import (
	"fmt"
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

	err := e.Start(":3000")
	if err != nil {
		fmt.Println("Error starting server: ", err.Error())
	}
}

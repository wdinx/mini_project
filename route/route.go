package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"mini_project/controller"
	"mini_project/repository"
	"mini_project/service"
)

func InitRoute(db *gorm.DB, e *echo.Echo) {

	adminRepository := repository.NewAdminRepository(db)

	adminService := service.NewAdminRepository(adminRepository)

	adminController := controller.NewAdminController(adminService)

	eAdmin := e.Group("/admin")
	eAdmin.POST("/login", adminController.Login)
	eAdmin.POST("/register", adminController.Register)

}

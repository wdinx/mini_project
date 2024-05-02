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
	userRepository := repository.NewUserRepository(db)

	adminService := service.NewAdminRepository(adminRepository)
	userService := service.NewUserService(userRepository)

	adminController := controller.NewAdminController(adminService)
	userController := controller.NewUserController(userService)
	fileController := controller.NewFileController()

	e.GET("/image/:image", fileController.ShowFile)

	eAdmin := e.Group("/v1/admin")
	eAdmin.POST("/login", adminController.Login)
	eAdmin.POST("/register", adminController.Register)

	eUser := e.Group("/v1/user")
	eUser.POST("/login", userController.Login)
	eUser.POST("/register", userController.Register)

}

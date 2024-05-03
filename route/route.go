package route

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"mini_project/constant"
	"mini_project/controller"
	"mini_project/repository"
	"mini_project/service"
)

func InitRoute(db *gorm.DB, e *echo.Echo, validate *validator.Validate) {

	adminRepository := repository.NewAdminRepository(db)
	userRepository := repository.NewUserRepository(db)

	adminService := service.NewAdminRepository(adminRepository, validate)
	userService := service.NewUserService(userRepository, validate)

	adminController := controller.NewAdminController(adminService)
	userController := controller.NewUserController(userService)
	fileController := controller.NewFileController()

	e.GET("/image/:image", fileController.ShowFile)

	eAdmin := e.Group("/v1/admin")
	eAdmin.POST("/login", adminController.Login)
	eAdmin.POST("/register", adminController.Register)
	eAdmin.Use(echojwt.JWT([]byte(constant.ADMIN_SECRET_JWT)))

	eUser := e.Group("/v1/user")
	eUser.POST("/login", userController.Login)
	eUser.POST("/register", userController.Register)
	eUser.Use(echojwt.JWT([]byte(constant.USER_SECRET_JWT)))

}

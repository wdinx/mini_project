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
	touristAttractionTypeRepository := repository.NewTouristAttractionTypeRepository(db)

	adminService := service.NewAdminRepository(adminRepository, validate)
	userService := service.NewUserService(userRepository, validate)
	touristAttractionTypeService := service.NewTouristAttractionTypeService(touristAttractionTypeRepository, validate)

	adminController := controller.NewAdminController(adminService)
	userController := controller.NewUserController(userService)
	fileController := controller.NewFileController()
	touristAttractionTypeController := controller.NewTouristAttractionTypeController(touristAttractionTypeService)

	e.GET("/image/:image", fileController.ShowFile)

	eA := e.Group("/v1/admin")
	eA.POST("/login", adminController.Login)
	eA.POST("/register", adminController.Register)

	eU := e.Group("/v1/user")
	eU.POST("/login", userController.Login)
	eU.POST("/register", userController.Register)

	eAdmin := e.Group("/v1/admin")
	eAdmin.Use(echojwt.JWT([]byte(constant.ADMIN_SECRET_JWT)))
	eAdmin.GET("/tourist-attraction-types", touristAttractionTypeController.GetAll)
	eAdmin.DELETE("/tourist-attraction-types/:id", touristAttractionTypeController.Delete)
	eAdmin.POST("/tourist-attraction-types", touristAttractionTypeController.Create)
	eAdmin.PUT("/tourist-attraction-types/:id", touristAttractionTypeController.Update)

	eUser := e.Group("/v1/user")
	eUser.Use(echojwt.JWT([]byte(constant.USER_SECRET_JWT)))

}

package route

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"mini_project/config"
	"mini_project/constant"
	"mini_project/controller"
	"mini_project/repository"
	"mini_project/service"
)

func InitRoute(db *gorm.DB, e *echo.Echo, validate *validator.Validate, config *config.Config) {

	adminRepository := repository.NewAdminRepository(db)
	userRepository := repository.NewUserRepository(db)
	touristAttractionTypeRepository := repository.NewTouristAttractionTypeRepository(db)
	touristAttractionRepository := repository.NewTouristAttractionRepository(db)
	paymentRepository := repository.NewPaymentRepository(db)

	adminService := service.NewAdminRepository(adminRepository, validate)
	userService := service.NewUserService(userRepository, validate)
	touristAttractionTypeService := service.NewTouristAttractionTypeService(touristAttractionTypeRepository, validate)
	touristAttractionService := service.NewTouristAttractionService(touristAttractionRepository, validate)
	midtransService := service.NewMidtransService(config)
	paymentService := service.NewPaymentService(paymentRepository, midtransService, touristAttractionRepository)

	adminController := controller.NewAdminController(adminService)
	userController := controller.NewUserController(userService)
	fileController := controller.NewFileController()
	touristAttractionTypeController := controller.NewTouristAttractionTypeController(touristAttractionTypeService)
	touristAttractionController := controller.NewTouristAttractionController(touristAttractionService)
	midtransController := controller.NewMidtransController(midtransService, paymentService)
	paymentController := controller.NewPaymentController(paymentService)

	e.GET("/image/:image", fileController.ShowFile)

	eA := e.Group("/v1/admin")
	eA.POST("/login", adminController.Login)
	eA.POST("/register", adminController.Register)

	eU := e.Group("/v1/user")
	eU.POST("/login", userController.Login)
	eU.POST("/register", userController.Register)
	// Route for Midtrans
	eU.POST("/midtrans/payment-callback", midtransController.PaymentHandler)

	eAdmin := e.Group("/v1/admin")
	eAdmin.Use(echojwt.JWT([]byte(constant.ADMIN_SECRET_JWT)))

	// Route for Tourist Attraction Type
	eAdmin.GET("/tourist-attraction-types", touristAttractionTypeController.GetAll)
	eAdmin.DELETE("/tourist-attraction-types/:id", touristAttractionTypeController.Delete)
	eAdmin.POST("/tourist-attraction-types", touristAttractionTypeController.Create)
	eAdmin.PUT("/tourist-attraction-types/:id", touristAttractionTypeController.Update)

	// Route for Tourist Attraction
	eAdmin.POST("/tourist-attractions", touristAttractionController.Create)
	eAdmin.PUT("/tourist-attractions/:id", touristAttractionController.Update)
	eAdmin.GET("/tourist-attractions", touristAttractionController.GetAll)
	eAdmin.PUT("/tourist-attractions/:id/balance", touristAttractionController.UpdateBalanceById)

	eUser := e.Group("/v1/user")
	eUser.Use(echojwt.JWT([]byte(constant.USER_SECRET_JWT)))

	// Route for Payment
	eUser.POST("/payments/initialize", paymentController.InitializePayment)

}

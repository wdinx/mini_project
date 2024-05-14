package route

import (
	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
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
	transactionRepository := repository.NewTransactionRepository(db)
	ticketRepository := repository.NewTicketRepository(db)
	imageRepository := repository.NewImageRepository(config.DigitalOceanSpaces)

	adminService := service.NewAdminService(adminRepository, validate)
	ImageService := service.NewImageService(imageRepository)
	userService := service.NewUserService(userRepository, ImageService, validate)
	touristAttractionTypeService := service.NewTouristAttractionTypeService(touristAttractionTypeRepository, validate)
	touristAttractionService := service.NewTouristAttractionService(touristAttractionRepository, ImageService, validate)
	midtransService := service.NewMidtransService(config)
	paymentService := service.NewPaymentService(paymentRepository, midtransService, touristAttractionService, ticketRepository, transactionRepository)
	transactionService := service.NewTransactionService(transactionRepository, touristAttractionRepository, paymentService, validate)
	ticketService := service.NewTicketService(ticketRepository, validate)

	adminController := controller.NewAdminController(adminService)
	userController := controller.NewUserController(userService)
	fileController := controller.NewFileController()
	touristAttractionTypeController := controller.NewTouristAttractionTypeController(touristAttractionTypeService)
	touristAttractionController := controller.NewTouristAttractionController(touristAttractionService)
	midtransController := controller.NewMidtransController(midtransService, paymentService)
	transactionController := controller.NewTransactionController(transactionService)
	ticketController := controller.NewTicketController(ticketService)

	e.GET("/image/:image", fileController.ShowFile)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"status":  200,
			"message": "succesfully connected",
		})
	})

	// Route For Admin Login and Register
	eA := e.Group("/v1/admin")
	eA.POST("/login", adminController.Login)
	eA.POST("/register", adminController.Register)

	// Route For User Login and Register
	eU := e.Group("/v1/user")
	eU.POST("/login", userController.Login)
	eU.POST("/register", userController.Register)

	// Route for Midtrans
	eU.POST("/midtrans/payment-callback", midtransController.PaymentHandler)

	eAdmin := e.Group("/v1/admin")
	eAdmin.Use(echojwt.JWT([]byte(constant.ADMIN_SECRET_JWT)))

	// Route for Tourist Attraction Type
	eAdmin.DELETE("/tourist-attraction-types/:id", touristAttractionTypeController.Delete)
	eAdmin.POST("/tourist-attraction-types", touristAttractionTypeController.Create)
	eAdmin.PUT("/tourist-attraction-types/:id", touristAttractionTypeController.Update)

	// Route for Tourist Attraction
	eAdmin.POST("/tourist-attractions", touristAttractionController.Create)
	eAdmin.PUT("/tourist-attractions/:id", touristAttractionController.Update)
	eAdmin.PUT("/tourist-attractions/:id/balance", touristAttractionController.UpdateBalanceById)

	// Route for Ticket
	eAdmin.GET("/tourist-attraction/:id/tickets", ticketController.FindByTouristAttractionID)
	eAdmin.GET("/ticket/:id", ticketController.FindByID)

	eUser := e.Group("/v1/user")
	eUser.Use(echojwt.JWT([]byte(constant.USER_SECRET_JWT)))

	// Route for Transaction
	eUser.POST("/transaction/initialize", transactionController.InitializeTransaction)
	eUser.GET("/transactions", transactionController.GetByUserID)
	eUser.GET("/transaction/:id", transactionController.GetByID)

	// Route For Ticket
	eUser.GET("/ticket/:id", ticketController.FindByID)
	eUser.GET("/tickets", ticketController.FindByUserID)

	// Route For All User
	eAllUser := e.Group("/v1")
	eAllUser.GET("/tourist-attractions", touristAttractionController.GetAll)
	eAllUser.GET("/tourist-attraction-types", touristAttractionTypeController.GetAll)
	eAllUser.GET("/tourist-attraction/:id", touristAttractionController.GetById)

}

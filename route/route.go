package route

import (
	"mini_project/config"
	"mini_project/constant"
	"mini_project/controller"
	"mini_project/repository"
	"mini_project/service"

	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	transactionService := service.NewTransactionService(transactionRepository, touristAttractionRepository, userRepository, paymentService, validate)
	ticketService := service.NewTicketService(ticketRepository, userRepository, validate)
	chatbotService := service.NewChatBotService()

	adminController := controller.NewAdminController(adminService)
	userController := controller.NewUserController(userService)
	touristAttractionTypeController := controller.NewTouristAttractionTypeController(touristAttractionTypeService)
	touristAttractionController := controller.NewTouristAttractionController(touristAttractionService)
	midtransController := controller.NewMidtransController(midtransService, paymentService)
	transactionController := controller.NewTransactionController(transactionService)
	ticketController := controller.NewTicketController(ticketService)
	chatbotController := controller.NewChatBotController(chatbotService)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"status":  200,
			"message": "succesfully connected",
		})
	})

	// Route Group
	eAdminAuth := e.Group("/v1/admin") // Route Group for Admin Login and Register
	eUserAuth := e.Group("/v1/user")   // Route Group for User Login and Register
	eAdmin := e.Group("/v1/admin")     // Route Group for Admin With Middleware
	eAdmin.Use(echojwt.JWT([]byte(constant.ADMIN_SECRET_JWT)))
	eUser := e.Group("/v1/user") // Route Group for User With Middleware
	eUser.Use(echojwt.JWT([]byte(constant.USER_SECRET_JWT)))
	eAllUser := e.Group("/v1") // Route Group for All User

	// Route For Admin Login and Register
	eAdminAuth.POST("/login", adminController.Login)
	eAdminAuth.POST("/register", adminController.Register)

	// Route For User Login and Register
	eUserAuth.POST("/login", userController.Login)
	eUserAuth.POST("/register", userController.Register)

	// Route for Midtrans
	eUserAuth.POST("/midtrans/payment-callback", midtransController.PaymentHandler)

	// Route for Tourist Attraction Type
	eAdmin.DELETE("/tourist-attraction-types/:id", touristAttractionTypeController.Delete)
	eAdmin.POST("/tourist-attraction-types", touristAttractionTypeController.Create)
	eAdmin.PUT("/tourist-attraction-types/:id", touristAttractionTypeController.Update)
	eAllUser.GET("/tourist-attraction-types", touristAttractionTypeController.GetAll)
	eAllUser.GET("/tourist-attraction-type/:id", touristAttractionTypeController.FindByID)

	// Route for Tourist Attraction
	eAdmin.POST("/tourist-attractions", touristAttractionController.Create)
	eAdmin.PUT("/tourist-attractions/:id", touristAttractionController.Update)
	eAllUser.GET("/tourist-attractions", touristAttractionController.GetAll)
	eAllUser.GET("/tourist-attraction/:id", touristAttractionController.GetById)

	// Route for Ticket
	eAdmin.GET("/tourist-attraction/:id/tickets", ticketController.FindByTouristAttractionID)
	eAdmin.GET("/ticket/:id", ticketController.FindByID)
	eUser.GET("/ticket/:id", ticketController.FindByID)
	eUser.GET("/tickets", ticketController.FindByUserID)

	// Route for Transaction
	eUser.POST("/transaction/initialize", transactionController.InitializeTransaction)
	eUser.GET("/transactions", transactionController.GetByUserID)
	eUser.GET("/transaction/:id", transactionController.GetByID)

	// Route for Chatbot
	e.GET("/chat", chatbotController.ChatBot)
}

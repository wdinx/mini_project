package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"mini_project/repository"
	"mini_project/service"
)

func InitRoute(db *gorm.DB, e *echo.Echo) {

	adminRepository := repository.NewAdminRepository(db)

	adminService := service.NewAdminRepository(adminRepository)

}

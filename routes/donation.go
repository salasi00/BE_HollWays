package routes

import (
	"holyways/handlers"
	"holyways/pkg/middleware"
	"holyways/pkg/mysql"
	"holyways/repositories"

	"github.com/labstack/echo/v4"
)

func DonationRoutes(e *echo.Group) {
	DonationRepository := repositories.RepositoryDonation(mysql.DB)
	h := handlers.HandlerDonation(DonationRepository)

	e.GET("/donations", h.FindDonation)
	e.GET("/donation/:id", h.GetDonation)
	e.POST("/donation", middleware.Auth(middleware.UploadFile(h.CreateDonation)))
	e.GET("/donation-user", middleware.Auth(h.GetDonationUser))
}

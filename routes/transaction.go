package routes

import (
	"holyways/handlers"
	"holyways/pkg/middleware"
	"holyways/pkg/mysql"
	"holyways/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)

	h := handlers.HandlerTransaction(TransactionRepository)

	e.GET("/transactions", h.FindTransaction)
	e.GET("/transaction", middleware.Auth(h.GetTransaction))
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	// e.GET("transaction-donation/:id", h.GetTransactionByDonation)
	e.POST("/notification", h.Notification)
	e.GET("/transaction-user", middleware.Auth(h.FindTransactionUser))
	e.GET("/transaction-donation", h.FindTransactionDonation)

}

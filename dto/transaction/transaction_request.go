package transactiondto

import "holyways/models"

type TransactionRequest struct {
	Total      int                             `json:"total" form:"total"`
	Status     string                          `json:"status" form:"status"`
	DonationID int                             `json:"donation_id" form:"donation_id"`
	Donation   models.DonationResponse         `json:"donation"`
	UserID     int                             `json:"user_id"`
	User       models.UsersTransactionResponse `json:"user"`
}

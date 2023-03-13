package transactiondto

import "holyways/models"

type TransactionUserResponse struct {
	ID        int                             `json:"id"`
	UserID    int                             `json:"user_id"`
	User      models.UsersTransactionResponse `json:"user"`
	CreatedAt string                          `json:"donate_at"`
	Total     int                             `json:"total"`
}

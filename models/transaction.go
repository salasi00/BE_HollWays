package models

type Transaction struct {
	ID         int                  `json:"id"`
	CreatedAt  string               `json:"donate_at"`
	Total      int                  `json:"total"`
	Status     string               `json:"status"`
	DonationID int                  `json:"donation_id" form:"donation_id"`
	Donation   DonationResponse     `json:"donation" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID     int                  `json:"-"`
	User       UsersProfileResponse `json:"user"`
}

type TransactionUser struct {
	ID        int                      `json:"id"`
	CreatedAt string                   `json:"donate_at"`
	Total     int                      `json:"total"`
	UserID    int                      `json:"-"`
	User      UsersTransactionResponse `json:"user"`
}

func (TransactionUser) TableName() string {
	return "transactions"
}

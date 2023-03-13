package models

import "time"

type User struct {
	ID        int                    `json:"id"`
	Fullname  string                 `json:"fullname" gorm:"type: varchar(255)"`
	Email     string                 `json:"email" gorm:"type: varchar(255)"`
	Password  string                 `json:"password" type:"varchar(255)"`
	Phone     string                 `json:"phone" gorm:"type: varchar(255)"`
	Image     string                 `json:"image" gorm:"type: varchar(255)"`
	Address   string                 `json:"address" gorm:"type: varchar(255)"`
	Donation  []DonationUserResponse `json:"donations"`
	CreatedAt time.Time              `json:"-"`
	UpdatedAt time.Time              `json:"-"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UsersTransactionResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
func (UsersTransactionResponse) TableName() string {
	return "users"
}

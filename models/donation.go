package models

import "time"

type Donation struct {
	ID            int                  `json:"id"`
	Title         string               `json:"title" gorm:"type: varchar(255)"`
	UserID        int                  `json:"user_id"`
	User          UsersProfileResponse `json:"user"`
	CurrentGoal   int                  `json:"current_goal"`
	Goal          int                  `json:"goal" gorm:"type: varchar(255)"`
	TotalDonation int                  `json:"total_donation" gorm:"type: varchar(255)"`
	Expired       string               `json:"expired" gorm:"type: varchar(255)"`
	Description   string               `json:"description" gorm:"type: varchar(255)"`
	Image         string               `json:"image"`
	CreatedAt     time.Time            `json:"-"`
	UpdatedAt     time.Time            `json:"-"`
}

type DonationResponse struct {
	ID          int                  `json:"id"`
	Title       string               `json:"title"`
	UserID      int                  `json:"-"`
	User        UsersProfileResponse `json:"user"`
	Goal        int                  `json:"goal"`
	CurrentGoal int                  `json:"current_goal"`
	Description string               `json:"description"`
	Image       string               `json:"image"`
}

type DonationUserResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CurrentGoal int    `json:"current_goal"`
	UserID      int    `json:"-"`
}

func (DonationUserResponse) TableName() string {
	return "donations"
}

func (DonationResponse) TableName() string {
	return "donations"
}

package authdto

type RegisterResponse struct {
	Email string `json:"email" gorm:"type: varchar(255)"`
	Token string `json:"token" gorm:"type: varchar(255)"`
}

type LoginResponse struct {
	Email string `json:"email" gorm:"type: varchar(255)"`
	Token string `json:"token" gorm:"type: varchar(255)"`
}

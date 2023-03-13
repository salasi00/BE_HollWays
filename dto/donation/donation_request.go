package donationdto

type DonationRequest struct {
	Title       string `json:"title" form:"title"`
	Goal        int    `json:"goal" form:"goal"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}

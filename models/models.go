package models

const (
	Success string = "Congratulations! You'r discount is 20%."
	Fail    string = "Oops! Code invalid or redeemed by another user."
)

type RedeemRequest struct {
	UserID       int64  `json:"user_id"`
	DiscountCode string `json:"discount_code"`
}

type RedeemResponse struct {
	Status string `json:"status"`
}

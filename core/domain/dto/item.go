package dto

import "time"

// It represents the data to be saved into the repository
// these fields are obtained  from the MeLi APIs
type ItemDto struct {
	Price       float32   `json:"price" binding:"required"`
	DateCreated time.Time `json:"date_created" binding:"required"`
	CategoryID  string    `json:"category_id" binding:"required"`
	CurrencyID  string    `json:"currency_id" binding:"required"`
	SellerID    uint      `json:"seller_id" binding:"required"`

	CategoryName        string `json:"name" binding:"required"`
	CurrencyDescription string `json:"description" binding:"required"`
	SellerNickName      string `json:"nickname" binding:"required"`
}

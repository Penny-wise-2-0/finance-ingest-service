package models

import (

	"gorm.io/gorm"
)





type Budget struct {
	gorm.Model
	UserID string
	Frequency string
    Category string
    Name string
    Amount string
}

type PlaidAccessToken struct {
	gorm.Model
	UserId string
	Token string
	ItemID string
}


    
type Transaction struct {
    gorm.Model
    UserID string
    AccountID string
    Amount float64
    Category string `gorm:"type:jsonb"`
    Date string
    Name string
    MerchantName string
    PaymentChannel string
    Pending bool
    TransactionID string
    TransactionType string
    ISOCurrencyCode string
	Address string
	City string
	
}


type Cursor struct {
	gorm.Model
	UserID string
	Cursor string
}
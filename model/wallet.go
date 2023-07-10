package model

import "time"

type Wallet struct {
	ID        string    `json:"id" gorm:"primary_key"`
	UserId    string    `json:"user_id"`
	User      *User     `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name      string    `json:"name" gorm:"name"`
	Balance   float32   `json:"balance" gorm:"balance"`
	Currency  string    `json:"currency" gorm:"currency"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

type WalletTransaction struct {
	ID          string    `json:"id" gorm:"primary_key"`
	WalletId    string    `json:"wallet_id"`
	Wallet      Wallet    `json:"wallet" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount      float32   `json:"amount" gorm:"amount"`
	Method      string    `json:"method" gorm:"method"`
	Currency    string    `json:"currency" gorm:"currency"`
	Status      string    `json:"status" gorm:"status"`
	Description string    `json:"description" gorm:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
	IsDeleted   bool      `json:"-" gorm:"is_deleted"`
}

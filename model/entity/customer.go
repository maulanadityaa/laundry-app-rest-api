package entity

import (
	"time"
)

type Customer struct {
	ID          string    `gorm:"primary_key" json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccountID   string    `json:"account_id"`

	Transaction []Transaction `gorm:"foreignKey:CustomerID" json:"transaction"`
}

func (c *Customer) TableName() string {
	return "mst_customer"
}

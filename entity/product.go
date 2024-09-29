package entity

import (
	"time"
)

type Product struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Price     uint      `json:"price"`
	Unit      string    `json:"unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	TransactionDetail []TransactionDetail `gorm:"foreignKey:ProductID" json:"transaction_detail"`
}

func (p *Product) TableName() string {
	return "mst_product"
}

package entity

import "time"

type Transaction struct {
	ID         string    `gorm:"primary_key" json:"id"`
	UserID     string    `json:"user_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	TotalPrice uint      `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionID" json:"transaction_detail"`
}

func (t *Transaction) TableName() string {
	return "trx_transaction"
}

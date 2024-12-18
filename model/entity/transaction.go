package entity

import "time"

type Transaction struct {
	ID         string    `gorm:"primary_key" json:"id"`
	CustomerID string    `json:"customer_id"`
	EmployeeID string    `json:"employee_id"`
	StartTime  time.Time `json:"start_time"`
	FinishTime time.Time `json:"finish_time"`
	TotalPrice uint      `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionID" json:"transaction_detail"`
}

func (t *Transaction) TableName() string {
	return "trx_transaction"
}

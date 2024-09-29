package entity

type TransactionDetail struct {
	ID            string `gorm:"primary_key" json:"id"`
	TransactionID string `json:"transaction_id"`
	ProductID     string `json:"product_id"`
	Quantity      uint   `json:"quantity"`
}

func (t *TransactionDetail) TableName() string {
	return "trx_transaction_detail"
}

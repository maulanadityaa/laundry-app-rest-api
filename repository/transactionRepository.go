package repository

import (
	"github.com/maulanadityaa/laundry-app-rest-api/model/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	AddTransaction(transaction entity.Transaction) (entity.Transaction, error)
	UpdateTransaction(transaction entity.Transaction) (entity.Transaction, error)
	GetAllTransaction(spec []func(db *gorm.DB) *gorm.DB, customerName, employeeName, startDate, endDate string) ([]entity.Transaction, string, error)
	GetTransactionByID(id string) (entity.Transaction, error)
}

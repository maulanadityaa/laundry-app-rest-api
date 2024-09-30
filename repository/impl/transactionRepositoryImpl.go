package impl

import (
	"github.com/maulanadityaa/laundry-app-rest-api/config"
	"github.com/maulanadityaa/laundry-app-rest-api/helper"
	"github.com/maulanadityaa/laundry-app-rest-api/model/entity"
	"gorm.io/gorm"
)

type TransactionRepository struct{}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (TransactionRepository) AddTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(&transaction); result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return entity.Transaction{}, err
	}

	return transaction, nil
}

func (TransactionRepository) UpdateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Save(&transaction); result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return entity.Transaction{}, err
	}

	return transaction, nil
}

func (TransactionRepository) GetAllTransaction(spec []func(db *gorm.DB) *gorm.DB, customerName, employeeName, startDate, endDate string) ([]entity.Transaction, string, error) {
	var transactions []entity.Transaction

	if customerName != "" {
		spec = append(spec, helper.SelectTransactionByCustomerName(customerName))
	}

	if employeeName != "" {
		spec = append(spec, helper.SelectTransactionByEmployeeName(employeeName))
	}

	if startDate != "" && endDate != "" {
		spec = append(spec, helper.SelectTransactionByTimePeriod(startDate, endDate))
	}

	db := config.DB.Model(&entity.Transaction{}).Scopes(spec[1:]...).Preload("TransactionDetails").Debug()
	totalRows := helper.GetTotalRows(db)
	err := db.Scopes(spec[0]).Find(&transactions).Error

	return transactions, totalRows, err
}

func (TransactionRepository) GetTransactionByID(id string) (entity.Transaction, error) {
	var transaction entity.Transaction

	if result := config.DB.Preload("TransactionDetails").Where("id = ?", id).First(&transaction); result.Error != nil {
		return entity.Transaction{}, result.Error
	}

	return transaction, nil
}

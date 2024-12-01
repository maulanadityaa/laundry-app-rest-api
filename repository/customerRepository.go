package repository

import (
	"maulanadityaa/laundry-app-rest-api/model/entity"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	AddCustomer(customer entity.Customer) (entity.Customer, error)
	UpdateCustomer(customer entity.Customer) (entity.Customer, error)
	GetCustomerByID(customerID string) (entity.Customer, error)
	GetCustomerByAccountID(accountID string) (entity.Customer, error)
	GetAllCustomer(spec []func(db *gorm.DB) *gorm.DB, name string) ([]entity.Customer, string, error)
}

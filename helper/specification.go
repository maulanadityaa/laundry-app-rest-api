package helper

import (
	"strings"

	"gorm.io/gorm"
)

func SelectByName(customerName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(customerName)+"%")
	}
}

func SelectTransactionByEmployeeName(employeeName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins("JOIN mst_employee ON mst_employee.id = transactions.employee_id").Where("LOWER(mst_employee.name) LIKE ?", "%"+strings.ToLower(employeeName)+"%")
	}
}

func SelectTransactionByCustomerName(customerName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins("JOIN mst_customer ON mst_customer.id = transactions.customer_id").Where("LOWER(mst_customer.name) LIKE ?", "%"+strings.ToLower(customerName)+"%")
	}
}

func SelectTransactionByTimePeriod(startDate, endDate string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("transactions.created_at BETWEEN ? AND ?", startDate, endDate)
	}
}

func SelectTransactionByStatus(status string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("transactions.status = ?", status)
	}
}

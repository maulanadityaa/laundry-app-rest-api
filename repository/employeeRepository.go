package repository

import (
	"github.com/maulanadityaa/laundry-app-rest-api/model/entity"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	AddEmployee(employee entity.Employee) (entity.Employee, error)
	UpdateEmployee(employee entity.Employee) (entity.Employee, error)
	GetEmployeeByID(employeeID string) (entity.Employee, error)
	GetEmployeeByAccountID(accountID string) (entity.Employee, error)
	GetAllEmployee(spec []func(db *gorm.DB) *gorm.DB, name string) ([]entity.Employee, string, error)
}

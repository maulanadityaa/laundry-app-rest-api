package impl

import (
	"maulanadityaa/laundry-app-rest-api/config"
	"maulanadityaa/laundry-app-rest-api/helper"
	"maulanadityaa/laundry-app-rest-api/model/entity"

	"gorm.io/gorm"
)

type EmployeeRepository struct{}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{}
}

func (EmployeeRepository) AddEmployee(employee entity.Employee) (entity.Employee, error) {
	if result := config.DB.Create(&employee); result.Error != nil {
		return entity.Employee{}, result.Error
	}

	return employee, nil
}

func (EmployeeRepository) UpdateEmployee(employee entity.Employee) (entity.Employee, error) {
	if result := config.DB.Save(&employee); result.Error != nil {
		return entity.Employee{}, result.Error
	}

	return employee, nil
}

func (EmployeeRepository) GetEmployeeByID(employeeID string) (entity.Employee, error) {
	var employee entity.Employee

	if err := config.DB.Where("id = ?", employeeID).First(&employee).Error; err != nil {
		return entity.Employee{}, err
	}

	return employee, nil
}

func (EmployeeRepository) GetEmployeeByAccountID(accountID string) (entity.Employee, error) {
	var employee entity.Employee

	if err := config.DB.Where("account_id = ?", accountID).First(&employee).Error; err != nil {
		return entity.Employee{}, err
	}

	return employee, nil
}

func (EmployeeRepository) GetAllEmployee(spec []func(db *gorm.DB) *gorm.DB, name string) ([]entity.Employee, string, error) {
	var employees []entity.Employee

	if name != "" {
		spec = append(spec, helper.SelectByName(name))
	}

	db := config.DB.Model(&entity.Employee{}).Scopes(spec[1:]...)
	totalRows := helper.GetTotalRows(db)
	err := db.Scopes(spec[0]).Find(&employees).Error

	return employees, totalRows, err
}

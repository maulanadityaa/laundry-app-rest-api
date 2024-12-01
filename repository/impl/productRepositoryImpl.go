package impl

import (
	"maulanadityaa/laundry-app-rest-api/config"
	"maulanadityaa/laundry-app-rest-api/helper"
	"maulanadityaa/laundry-app-rest-api/model/entity"

	"gorm.io/gorm"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (ProductRepository) AddProduct(product entity.Product) (entity.Product, error) {
	if result := config.DB.Create(&product); result.Error != nil {
		return entity.Product{}, result.Error
	}

	return product, nil
}

func (ProductRepository) UpdateProduct(product entity.Product) (entity.Product, error) {
	if result := config.DB.Save(&product); result.Error != nil {
		return entity.Product{}, result.Error
	}

	return product, nil
}

func (ProductRepository) DeleteProduct(id string) error {
	if result := config.DB.Delete(&entity.Product{}, "id = ?", id); result.Error != nil {
		return result.Error
	}

	return nil
}

func (ProductRepository) GetAllProduct(spec []func(db *gorm.DB) *gorm.DB) ([]entity.Product, string, error) {
	var products []entity.Product

	db := config.DB.Model(&entity.Product{}).Scopes(spec[1:]...)
	totalRows := helper.GetTotalRows(db)
	err := db.Scopes(spec[0]).Find(&products).Error

	return products, totalRows, err
}

func (ProductRepository) GetProductByID(id string) (entity.Product, error) {
	var product entity.Product

	if result := config.DB.Where("id = ?", id).First(&product); result.Error != nil {
		return entity.Product{}, result.Error
	}

	return product, nil
}

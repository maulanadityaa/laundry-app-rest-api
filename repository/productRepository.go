package repository

import (
	"github.com/maulanadityaa/laundry-app-rest-api/model/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	AddProduct(product entity.Product) (entity.Product, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
	DeleteProduct(id string) error
	GetAllProduct(spec []func(db *gorm.DB) *gorm.DB) ([]entity.Product, string, error)
	GetProductByID(id string) (entity.Product, error)
}

package service

import (
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/request"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/response"
)

type ProductService interface {
	AddProduct(req request.ProductRequest) (response.ProductResponse, error)
	UpdateProduct(req request.ProductUpdateRequest) (response.ProductResponse, error)
	DeleteProduct(id string) error
	GetAllProduct(paging, rowsPerPage, name string) ([]response.ProductResponse, string, string, error)
	GetProductByID(id string) (response.ProductResponse, error)
}

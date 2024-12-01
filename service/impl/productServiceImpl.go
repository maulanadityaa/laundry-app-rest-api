package impl

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"maulanadityaa/laundry-app-rest-api/helper"
	"maulanadityaa/laundry-app-rest-api/model/dto/request"
	"maulanadityaa/laundry-app-rest-api/model/dto/response"
	"maulanadityaa/laundry-app-rest-api/model/entity"
	"maulanadityaa/laundry-app-rest-api/repository"
	"maulanadityaa/laundry-app-rest-api/repository/impl"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductService struct{}

var productRepository repository.ProductRepository = impl.NewProductRepository()

func NewProductService() *ProductService {
	return &ProductService{}
}

func (ProductService) AddProduct(req request.ProductRequest) (response.ProductResponse, error) {
	newProduct := entity.Product{
		ID:        uuid.NewString(),
		Name:      req.Name,
		Price:     req.Price,
		Unit:      req.Unit,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := productRepository.AddProduct(newProduct)
	if err != nil {
		fmt.Println(err)
		return response.ProductResponse{}, err
	}

	return response.ProductResponse{
		ID:        result.ID,
		Name:      result.Name,
		Price:     result.Price,
		Unit:      result.Unit,
		CreatedAt: result.CreatedAt.String(),
		UpdatedAt: result.UpdatedAt.String(),
	}, nil
}

func (ProductService) UpdateProduct(req request.ProductUpdateRequest) (response.ProductResponse, error) {
	product, err := productRepository.GetProductByID(req.ID)
	if err != nil {
		fmt.Println(err)
		return response.ProductResponse{}, err
	}

	product.Name = req.Name
	product.Price = req.Price
	product.Unit = req.Unit
	product.UpdatedAt = time.Now()

	result, err := productRepository.UpdateProduct(product)
	if err != nil {
		fmt.Println(err)
		return response.ProductResponse{}, err
	}

	return response.ProductResponse{
		ID:        result.ID,
		Name:      result.Name,
		Price:     result.Price,
		Unit:      result.Unit,
		CreatedAt: result.CreatedAt.String(),
		UpdatedAt: result.UpdatedAt.String(),
	}, nil
}

func (ProductService) DeleteProduct(id string) error {
	err := productRepository.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}

func (ProductService) GetAllProduct(paging, rowsPerPage, name string) ([]response.ProductResponse, string, string, error) {
	pagingInt, err := strconv.Atoi(paging)
	if err != nil {
		return nil, "0", "0", errors.New("invalid query parameter")
	}

	rowsPerPageInt, err := strconv.Atoi(rowsPerPage)
	if err != nil {
		return nil, "0", "0", errors.New("invalid query parameter")
	}

	var spec []func(db *gorm.DB) *gorm.DB
	spec = append(spec, helper.Paginate(pagingInt, rowsPerPageInt))

	if name != "" {
		spec = append(spec, helper.SelectByName(name))
	}

	products, totalRows, err := productRepository.GetAllProduct(spec)
	if err != nil {
		fmt.Println(err)
		return nil, "0", "0", err
	}

	productResponses := make([]response.ProductResponse, 0)
	for _, product := range products {
		productResponses = append(productResponses, response.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			Price:     product.Price,
			Unit:      product.Unit,
			CreatedAt: product.CreatedAt.String(),
			UpdatedAt: product.UpdatedAt.String(),
		})
	}

	totalPage := helper.GetTotalPage(totalRows, rowsPerPageInt)

	return productResponses, totalRows, strconv.Itoa(totalPage), nil
}

func (ProductService) GetProductByID(id string) (response.ProductResponse, error) {
	product, err := productRepository.GetProductByID(id)
	if err != nil {
		fmt.Println(err)
		return response.ProductResponse{}, err
	}

	return response.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Unit:      product.Unit,
		CreatedAt: product.CreatedAt.String(),
		UpdatedAt: product.UpdatedAt.String(),
	}, nil
}

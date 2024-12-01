package impl

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"maulanadityaa/laundry-app-rest-api/helper"
	"maulanadityaa/laundry-app-rest-api/model/dto/request"
	"maulanadityaa/laundry-app-rest-api/model/dto/response"
	"maulanadityaa/laundry-app-rest-api/model/entity"
	"maulanadityaa/laundry-app-rest-api/repository"
	"maulanadityaa/laundry-app-rest-api/repository/impl"
	"maulanadityaa/laundry-app-rest-api/service"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionService struct{}

var transactionRepository repository.TransactionRepository = impl.NewTransactionRepository()
var productService service.ProductService = NewProductService()

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (TransactionService) AddTransaction(req request.TransactionRequest) (response.TransactionResponse, error) {
	newTransaction := entity.Transaction{
		ID:         uuid.NewString(),
		CustomerID: req.CustomerID,
		EmployeeID: req.EmployeeID,
		StartTime:  req.StartTime,
		FinishTime: req.FinishTime,
		Status:     "PENDING",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	for _, trxReq := range req.ProductTransactionRequest {
		productResponse, _ := productService.GetProductByID(trxReq.ProductID)

		newTransaction.TransactionDetails = append(newTransaction.TransactionDetails, entity.TransactionDetail{
			ID:            uuid.NewString(),
			TransactionID: newTransaction.ID,
			ProductID:     trxReq.ProductID,
			Quantity:      trxReq.Quantity,
		})

		newTransaction.TotalPrice += productResponse.Price * trxReq.Quantity
	}

	customerResponse, _ := customerService.GetCustomerByID(req.CustomerID)
	employeeResponse, _ := employeeService.GetEmployeeByID(req.EmployeeID)

	result, err := transactionRepository.AddTransaction(newTransaction)
	if err != nil {
		return response.TransactionResponse{}, err
	}

	var transactionDetailsResponse []response.TransactionDetailResponse

	for _, trxDetail := range result.TransactionDetails {
		productResponse, _ := productService.GetProductByID(trxDetail.ProductID)

		transactionDetailsResponse = append(transactionDetailsResponse, response.TransactionDetailResponse{
			ID: trxDetail.ID,
			ProductTransactionResponse: response.ProductTransactionResponse{
				ID:       productResponse.ID,
				Name:     productResponse.Name,
				Price:    productResponse.Price,
				Unit:     productResponse.Unit,
				Quantity: trxDetail.Quantity,
			},
		})
	}

	return response.TransactionResponse{
		ID:                result.ID,
		CustomerResponse:  customerResponse,
		EmployeeResponse:  employeeResponse,
		StartTime:         result.StartTime.String(),
		FinishTime:        result.FinishTime.String(),
		TransactionDetail: transactionDetailsResponse,
		TotalPrice:        result.TotalPrice,
		Status:            result.Status,
		CreatedAt:         result.CreatedAt.String(),
		UpdatedAt:         result.UpdatedAt.String(),
	}, nil
}

func (TransactionService) UpdateTransaction(req request.TransactionUpdateRequest) (response.TransactionResponse, error) {
	var status string

	if strings.ToUpper(req.Status) == "ON PROCESS" {
		status = "ON PROCESS"
	} else if strings.ToUpper(req.Status) == "DONE" {
		status = "DONE"
	} else {
		return response.TransactionResponse{}, errors.New("invalid status")
	}

	updateTransaction, err := transactionRepository.GetTransactionByID(req.ID)
	if err != nil {
		return response.TransactionResponse{}, err
	}

	updateTransaction.Status = status

	result, err := transactionRepository.UpdateTransaction(updateTransaction)
	if err != nil {
		return response.TransactionResponse{}, err
	}

	var transactionDetailsResponse []response.TransactionDetailResponse

	for _, trxDetail := range result.TransactionDetails {
		productResponse, _ := productService.GetProductByID(trxDetail.ProductID)

		transactionDetailsResponse = append(transactionDetailsResponse, response.TransactionDetailResponse{
			ID: trxDetail.ID,
			ProductTransactionResponse: response.ProductTransactionResponse{
				ID:       productResponse.ID,
				Name:     productResponse.Name,
				Price:    productResponse.Price,
				Unit:     productResponse.Unit,
				Quantity: trxDetail.Quantity,
			},
		})
	}

	customerResponse, _ := customerService.GetCustomerByID(result.CustomerID)
	employeeResponse, _ := employeeService.GetEmployeeByID(result.EmployeeID)

	return response.TransactionResponse{
		ID:                result.ID,
		CustomerResponse:  customerResponse,
		EmployeeResponse:  employeeResponse,
		StartTime:         result.StartTime.String(),
		FinishTime:        result.FinishTime.String(),
		TransactionDetail: transactionDetailsResponse,
		TotalPrice:        result.TotalPrice,
		Status:            result.Status,
		CreatedAt:         result.CreatedAt.String(),
		UpdatedAt:         result.UpdatedAt.String(),
	}, nil
}

func (TransactionService) GetAllTransaction(paging, rowsPerPage, customerName, employeeName, startDate, endDate string) ([]response.TransactionResponse, string, string, error) {
	var transactions []response.TransactionResponse
	var totalRows string
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

	entityTransactions, totalRows, err := transactionRepository.GetAllTransaction(spec, customerName, employeeName, startDate, endDate)
	if err != nil {
		return nil, "", "", err

	}

	for _, transaction := range entityTransactions {
		var transactionDetailsResponse []response.TransactionDetailResponse

		for _, trxDetail := range transaction.TransactionDetails {
			productResponse, _ := productService.GetProductByID(trxDetail.ProductID)

			transactionDetailsResponse = append(transactionDetailsResponse, response.TransactionDetailResponse{
				ID: trxDetail.ID,
				ProductTransactionResponse: response.ProductTransactionResponse{
					ID:       productResponse.ID,
					Name:     productResponse.Name,
					Price:    productResponse.Price,
					Unit:     productResponse.Unit,
					Quantity: trxDetail.Quantity,
				},
			})
		}

		customerResponse, _ := customerService.GetCustomerByID(transaction.CustomerID)
		employeeResponse, _ := employeeService.GetEmployeeByID(transaction.EmployeeID)

		transactions = append(transactions, response.TransactionResponse{
			ID:                transaction.ID,
			CustomerResponse:  customerResponse,
			EmployeeResponse:  employeeResponse,
			StartTime:         transaction.StartTime.String(),
			FinishTime:        transaction.FinishTime.String(),
			TransactionDetail: transactionDetailsResponse,
			TotalPrice:        transaction.TotalPrice,
			Status:            transaction.Status,
			CreatedAt:         transaction.CreatedAt.String(),
			UpdatedAt:         transaction.UpdatedAt.String(),
		})
	}

	return transactions, totalRows, rowsPerPage, nil
}

func (TransactionService) GetTransactionByID(transactionID string) (response.TransactionResponse, error) {
	result, err := transactionRepository.GetTransactionByID(transactionID)
	if err != nil {
		return response.TransactionResponse{}, err
	}

	var transactionDetailsResponse []response.TransactionDetailResponse

	for _, trxDetail := range result.TransactionDetails {
		productResponse, _ := productService.GetProductByID(trxDetail.ProductID)

		transactionDetailsResponse = append(transactionDetailsResponse, response.TransactionDetailResponse{
			ID: trxDetail.ID,
			ProductTransactionResponse: response.ProductTransactionResponse{
				ID:       productResponse.ID,
				Name:     productResponse.Name,
				Price:    productResponse.Price,
				Unit:     productResponse.Unit,
				Quantity: trxDetail.Quantity,
			},
		})
	}

	customerResponse, _ := customerService.GetCustomerByID(result.CustomerID)
	employeeResponse, _ := employeeService.GetEmployeeByID(result.EmployeeID)

	return response.TransactionResponse{
		ID:                result.ID,
		CustomerResponse:  customerResponse,
		EmployeeResponse:  employeeResponse,
		StartTime:         result.StartTime.String(),
		FinishTime:        result.FinishTime.String(),
		TransactionDetail: transactionDetailsResponse,
		TotalPrice:        result.TotalPrice,
		Status:            result.Status,
		CreatedAt:         result.CreatedAt.String(),
		UpdatedAt:         result.UpdatedAt.String(),
	}, nil
}

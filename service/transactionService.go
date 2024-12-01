package service

import (
	"maulanadityaa/laundry-app-rest-api/model/dto/request"
	"maulanadityaa/laundry-app-rest-api/model/dto/response"
)

type TransactionService interface {
	AddTransaction(req request.TransactionRequest) (response.TransactionResponse, error)
	UpdateTransaction(req request.TransactionUpdateRequest) (response.TransactionResponse, error)
	GetAllTransaction(paging, rowsPerPage, customerName, employeeName, startDate, endDate string) ([]response.TransactionResponse, string, string, error)
	GetTransactionByID(id string) (response.TransactionResponse, error)
}

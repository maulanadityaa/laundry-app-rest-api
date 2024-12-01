package request

import "time"

type (
	ProductTransactionRequest struct {
		ProductID string `json:"productId" validate:"required" example:"ValidUUIDv4"`
		Quantity  uint   `json:"quantity" validate:"required" example:"1"`
	}

	TransactionRequest struct {
		CustomerID                string                      `json:"customerId" validate:"required" example:"ValidUUIDv4"`
		EmployeeID                string                      `json:"employeeId" validate:"required" example:"ValidUUIDv4"`
		StartTime                 time.Time                   `json:"startTime" validate:"required" example:"2021-01-01T00:00:00Z"`
		FinishTime                time.Time                   `json:"finishTime" validate:"required" example:"2021-01-01T00:00:00Z"`
		ProductTransactionRequest []ProductTransactionRequest `json:"products" validate:"required" example:"[{\"productId\":\"ValidUUIDv4\",\"quantity\":1}]"`
	}

	TransactionUpdateRequest struct {
		ID     string `json:"id" validate:"required" example:"ValidUUIDv4"`
		Status string `json:"status" validate:"required" example:"PENDING"`
	}
)

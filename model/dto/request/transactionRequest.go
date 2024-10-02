package request

import "time"

type (
	ProductTransactionRequest struct {
		ProductID string `json:"productId" validate:"required"`
		Quantity  uint   `json:"quantity" validate:"required"`
	}

	TransactionRequest struct {
		CustomerID                string                      `json:"customerId" validate:"required"`
		EmployeeID                string                      `json:"employeeId" validate:"required"`
		StartTime                 time.Time                   `json:"startTime" validate:"required"`
		FinishTime                time.Time                   `json:"finishTime" validate:"required"`
		ProductTransactionRequest []ProductTransactionRequest `json:"products" validate:"required"`
	}

	TransactionUpdateRequest struct {
		ID     string `json:"id" validate:"required"`
		Status string `json:"status" validate:"required"`
	}
)

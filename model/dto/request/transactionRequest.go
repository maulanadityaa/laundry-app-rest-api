package request

import "time"

type (
	ProductTransactionRequest struct {
		ProductID string `json:"productId" binding:"required"`
		Quantity  uint   `json:"quantity" binding:"required"`
	}

	TransactionRequest struct {
		CustomerID                string                      `json:"customerId" binding:"required"`
		EmployeeID                string                      `json:"employeeId" binding:"required"`
		StartTime                 time.Time                   `json:"startTime" binding:"required"`
		FinishTime                time.Time                   `json:"finishTime" binding:"required"`
		ProductTransactionRequest []ProductTransactionRequest `json:"products" binding:"required"`
	}

	TransactionUpdateRequest struct {
		ID     string `json:"id" binding:"required"`
		Status string `json:"status" binding:"required"`
	}
)

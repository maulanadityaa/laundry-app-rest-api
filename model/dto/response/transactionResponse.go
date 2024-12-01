package response

type (
	ProductTransactionResponse struct {
		ID       string `json:"id" example:"ValidUUIDv4"`
		Name     string `json:"name" example:"Product Name"`
		Price    uint   `json:"price" example:"10000"`
		Unit     string `json:"unit" example:"KG"`
		Quantity uint   `json:"quantity" example:"1"`
	}

	TransactionDetailResponse struct {
		ID                         string                     `json:"id" example:"ValidUUIDv4"`
		ProductTransactionResponse ProductTransactionResponse `json:"product" `
	}

	TransactionResponse struct {
		ID                string                      `json:"id" example:"ValidUUIDv4"`
		CustomerResponse  UserResponse                `json:"customer"`
		EmployeeResponse  UserResponse                `json:"employee"`
		TransactionDetail []TransactionDetailResponse `json:"transactionDetail"`
		StartTime         string                      `json:"startTime" example:"2021-01-01T00:00:00Z"`
		FinishTime        string                      `json:"finishTime" example:"2021-01-01T00:00:00Z"`
		TotalPrice        uint                        `json:"totalPrice" example:"10000"`
		Status            string                      `json:"status" example:"PENDING"`
		CreatedAt         string                      `json:"createdAt" example:"2021-01-01T00:00:00Z"`
		UpdatedAt         string                      `json:"updatedAt" example:"2021-01-01T00:00:00Z"`
	}
)

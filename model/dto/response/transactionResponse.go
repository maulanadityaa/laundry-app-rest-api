package response

type (
	ProductTransactionResponse struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Price    uint   `json:"price"`
		Unit     string `json:"unit"`
		Quantity uint   `json:"quantity"`
	}

	TransactionDetailResponse struct {
		ID                         string                     `json:"id"`
		ProductTransactionResponse ProductTransactionResponse `json:"product"`
	}

	TransactionResponse struct {
		ID                string                      `json:"id"`
		CustomerResponse  UserResponse                `json:"customer"`
		EmployeeResponse  UserResponse                `json:"employee"`
		TransactionDetail []TransactionDetailResponse `json:"transactionDetail"`
		StartTime         string                      `json:"startTime"`
		FinishTime        string                      `json:"finishTime"`
		TotalPrice        uint                        `json:"totalPrice"`
		Status            string                      `json:"status"`
		CreatedAt         string                      `json:"createdAt"`
		UpdatedAt         string                      `json:"updatedAt"`
	}
)

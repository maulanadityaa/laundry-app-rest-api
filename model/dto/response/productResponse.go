package response

type (
	ProductResponse struct {
		ID        string `json:"id" example:"ValidUUIDv4"`
		Name      string `json:"name" example:"Product Name"`
		Price     uint   `json:"price" example:"10000"`
		Unit      string `json:"unit" example:"KG"`
		CreatedAt string `json:"createdAt" example:"2021-01-01T00:00:00Z"`
		UpdatedAt string `json:"updatedAt" example:"2021-01-01T00:00:00Z"`
	}
)

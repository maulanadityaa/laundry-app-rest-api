package request

type (
	ProductRequest struct {
		Name  string `json:"name" validate:"required" example:"Product Name"`
		Price uint   `json:"price" validate:"required" example:"10000"`
		Unit  string `json:"unit" validate:"required" example:"KG"`
	}

	ProductUpdateRequest struct {
		ID    string `json:"id" validate:"required" example:"ValidUUIDv4"`
		Name  string `json:"name" validate:"required" example:"Product Name"`
		Price uint   `json:"price" validate:"required" example:"10000"`
		Unit  string `json:"unit" validate:"required" example:"KG"`
	}
)

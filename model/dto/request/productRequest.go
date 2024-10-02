package request

type (
	ProductRequest struct {
		Name  string `json:"name" validate:"required"`
		Price uint   `json:"price" validate:"required"`
		Unit  string `json:"unit" validate:"required"`
	}

	ProductUpdateRequest struct {
		ID    string `json:"id" validate:"required"`
		Name  string `json:"name" validate:"required"`
		Price uint   `json:"price" validate:"required"`
		Unit  string `json:"unit" validate:"required"`
	}
)

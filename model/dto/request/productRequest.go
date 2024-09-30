package request

type (
	ProductRequest struct {
		Name  string `json:"name" binding:"required"`
		Price uint   `json:"price" binding:"required"`
		Unit  string `json:"unit" binding:"required"`
	}

	ProductUpdateRequest struct {
		ID    string `json:"id" binding:"required"`
		Name  string `json:"name" binding:"required"`
		Price uint   `json:"price" binding:"required"`
		Unit  string `json:"unit" binding:"required"`
	}
)

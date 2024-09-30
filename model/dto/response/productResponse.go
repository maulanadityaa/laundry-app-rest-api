package response

type (
	ProductResponse struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Price     uint   `json:"price"`
		Unit      string `json:"unit"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}
)

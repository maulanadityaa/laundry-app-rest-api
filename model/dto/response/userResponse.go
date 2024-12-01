package response

type (
	UserResponse struct {
		ID          string `json:"id" example:"ValidUUIDv4"`
		Name        string `json:"name" example:"John Doe"`
		PhoneNumber string `json:"phoneNumber" example:"081234567890"`
		Address     string `json:"address" example:"Jl. Address No. 1"`
		CreatedAt   string `json:"createdAt" example:"2021-01-01T00:00:00Z"`
		UpdatedAt   string `json:"updatedAt" example:"2021-01-01T00:00:00Z"`
	}
)

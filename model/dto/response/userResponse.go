package response

type (
	UserResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phoneNumber"`
		Address     string `json:"address"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
	}
)

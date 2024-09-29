package request

type (
	UserRequest struct {
		Name        string `json:"name" binding:"required"`
		PhoneNumber string `json:"phoneNumber" binding:"required"`
		Address     string `json:"address" binding:"required"`
		AccountID   string `json:"accountId"`
	}

	UserUpdateRequest struct {
		ID          string `json:"id" binding:"required"`
		Name        string `json:"name" binding:"required"`
		PhoneNumber string `json:"phoneNumber" binding:"required"`
		Address     string `json:"address" binding:"required"`
	}
)

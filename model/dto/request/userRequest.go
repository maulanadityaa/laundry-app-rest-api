package request

type (
	UserRequest struct {
		Name        string `json:"name" validate:"required"`
		PhoneNumber string `json:"phoneNumber" validate:"required,uniquePhoneNumber,indonesianPhoneNumber"`
		Address     string `json:"address" validate:"required"`
		AccountID   string `json:"accountId"`
	}

	UserUpdateRequest struct {
		ID          string `json:"id" validate:"required"`
		Name        string `json:"name" validate:"required"`
		PhoneNumber string `json:"phoneNumber" validate:"required,uniquePhoneNumber,indonesianPhoneNumber"`
		Address     string `json:"address" validate:"required"`
	}
)

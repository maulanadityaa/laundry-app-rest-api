package request

type (
	UserRequest struct {
		Name        string `json:"name" validate:"required" example:"John Doe"`
		PhoneNumber string `json:"phoneNumber" validate:"required,uniquePhoneNumber,indonesianPhoneNumber" example:"081234567890"`
		Address     string `json:"address" validate:"required" example:"Jl. Address No. 1"`
		AccountID   string `json:"-"`
	}

	UserUpdateRequest struct {
		ID          string `json:"id" validate:"required" example:"ValidUUIDv4"`
		Name        string `json:"name" validate:"required" example:"John Doe"`
		PhoneNumber string `json:"phoneNumber" validate:"required,uniquePhoneNumber,indonesianPhoneNumber" example:"081234567890"`
		Address     string `json:"address" validate:"required" example:"Jl. Address No. 1"`
	}
)

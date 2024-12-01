package request

type (
	RegisterRequest struct {
		Email    string `json:"email" validate:"required,email,uniqueEmail" example:"johndoe@example.com"`
		Password string `json:"password" validate:"required" example:"12345678"`
		Role     string `json:"role" validate:"required,oneof=ROLE_CUSTOMER ROLE_EMPLOYEE" example:"ROLE_CUSTOMER"`
		UserRequest
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email" example:"johndoe@example.com"`
		Password string `json:"password" validate:"required" example:"12345678"`
	}
)

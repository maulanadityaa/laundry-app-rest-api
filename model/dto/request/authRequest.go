package request

type (
	RegisterRequest struct {
		Email    string `json:"email" validate:"required,email,uniqueEmail"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
		UserRequest
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)

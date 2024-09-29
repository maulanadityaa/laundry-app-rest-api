package response

type (
	LoginResponse struct {
		Token string `json:"token"`
	}

	RegisterResponse struct {
		ID           string       `json:"id"`
		Email        string       `json:"email"`
		Role         RoleResponse `json:"role"`
		UserResponse `json:"user"`
	}
)

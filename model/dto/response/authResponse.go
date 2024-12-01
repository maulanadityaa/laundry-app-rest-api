package response

type (
	LoginResponse struct {
		Token string `json:"token" example:"ValidJWTToken"`
	}

	RegisterResponse struct {
		ID           string       `json:"id" example:"ValidUUIDv4"`
		Email        string       `json:"email" example:"johndoe@example.com"`
		Role         RoleResponse `json:"role"`
		UserResponse `json:"user"`
	}
)

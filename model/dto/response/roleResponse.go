package response

type (
	RoleResponse struct {
		ID   string `json:"id" example:"ValidUUIDv4"`
		Name string `json:"name" example:"ROLE_CUSTOMER"`
	}
)

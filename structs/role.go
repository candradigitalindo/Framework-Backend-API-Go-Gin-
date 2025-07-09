package structs

type RoleCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type RoleResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

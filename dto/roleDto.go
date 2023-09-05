package dto

type RoleResponse struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
}

type CreateRoleRequest struct {
	Role string `json:"role" binding:"required"`
}

type UpdateRoleRequest struct {
	Role string `json:"role"`
}

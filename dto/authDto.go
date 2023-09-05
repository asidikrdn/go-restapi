package dto

import "github.com/google/uuid"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID       uuid.UUID    `json:"id,omitempty"`
	FullName string       `json:"fullname,omitempty"`
	Email    string       `json:"email,omitempty"`
	Role     RoleResponse `json:"role"`
	Token    string       `json:"token,omitempty"`
}

type RegisterRequest struct {
	FullName string `json:"fullname" form:"fullname" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Address  string `json:"address" form:"address" binding:"required"`
	RoleID   uint   `json:"roleId" form:"roleId"`
}

type VerifyEmailRequest struct {
	Email    string `json:"email" binding:"required"`
	OTPToken string `json:"otpToken" binding:"required"`
}

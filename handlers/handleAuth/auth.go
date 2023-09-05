package handlerAuth

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/repositories"
)

type handlerAuth struct {
	UserRepository repositories.UserRepository
}

func HandlerAuth(userRepository repositories.UserRepository) *handlerAuth {
	return &handlerAuth{userRepository}
}

func convertRegisterResponse(user *models.MstUser) *dto.UserResponse {
	return &dto.UserResponse{
		ID:              user.ID,
		FullName:        user.FullName,
		Email:           user.Email,
		IsEmailVerified: user.IsEmailVerified,
		Phone:           user.Phone,
		IsPhoneVerified: user.IsPhoneVerified,
		Address:         user.Address,
		Image:           user.Image,
		Role: dto.RoleResponse{
			ID:   user.Role.ID,
			Role: user.Role.Role,
		},
	}
}

func convertLoginResponse(user *models.MstUser, token string) *dto.LoginResponse {
	return &dto.LoginResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Role: dto.RoleResponse{
			ID:   user.Role.ID,
			Role: user.Role.Role,
		},
		Token: token,
	}
}

package handlerUser

import (
	"go-restapi-boilerplate/db/models"
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/repositories"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(userRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{userRepository}
}

func convertUserResponse(user *models.MstUser) *dto.UserResponse {
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

func convertMultipleUserResponse(user *[]models.MstUser) *[]dto.UserResponse {
	var users []dto.UserResponse

	for _, u := range *user {
		users = append(users, *convertUserResponse(&u))
	}

	return &users
}

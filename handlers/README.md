# HANDLERS

This folder contains all functions for handling any request from client and serve the response.

example :
`user.go`

```go
package handlers

import (
 "go-restapi-boilerplate/dto"
 "go-restapi-boilerplate/db/models"
 "go-restapi-boilerplate/repositories"
)

type handlerUser struct {
 UserRepository repositories.UserRepository
}

func HandlerUser(userRepository repositories.UserRepository) *handlerUser {
 return &handlerUser{userRepository}
}

func (h *handlerUser) FindAllUsers(c *gin.Context) {
 users, err := h.UserRepository.FindAllUsers()
 if err != nil {
  response := dto.Result{
   Status:  http.StatusInternalServerError,
   Message: err.Error(),
   Data:    nil,
  }
  c.JSON(http.StatusInternalServerError, response)
  return
 }

 response := dto.Result{
  Status:  http.StatusOK,
  Message: "Success",
  Data:    convertMultipleUserResponse(users),
 }
 c.JSON(http.StatusOK, response)
}

// convert response
func convertUserResponse(user *models.User) *dto.CustomerResponse {
 return &dto.UserResponse{
  ID:               user.ID,
  FullName:         user.FullName,
  Email:            user.Email,
  }
}
func convertMultipleUserResponse(users *[]models.User) *[]dto.UserResponse {
 var userResponse []dto.UserResponse

 for _, user := range *users {
  userResponse = append(userResponse, *convertCustomerResponse(&user))
 }

 return &userResponse
}
```

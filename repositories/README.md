# REPOSITORIES

This folder contains many use-cases for comunicate operations between this application and the database. Any function on repository didn't have manipulation logic, main function of repository just to get data from database and return it to any handlers that need it or vice versa.

example :
`user.go`

```go
package repositories

import (
 "go-restapi-boilerplate/models"
)

type UserRepository interface {
 FindAllUsers() (*[]models.User, error)
 }

func (r *repository) FindAllUsers() (*[]models.User, error) {
 var users []models.User
 err := r.db.Find(&users).Error
 return &users, err
}
```

# DATA TRANSFER OBJECT

This folder contains all struct that will be used for request and response processing.

example :
`user.go`

```go
package dto

import "go-restapi-boilerplate/models"

type UserRequest struct {
 FullName           string  `form:"fullname" json:"fullname"`
 Email              string  `form:"email" json:"email"`
 Password           string  `form:"password" json:"password"`
 }

type UserResponse struct {
 ID                 string  `form:"id" json:"id"`
 FullName           string  `form:"fullname" json:"fullname"`
 Email              string  `form:"email" json:"email"`
 }
```

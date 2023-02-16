# MODELS

This folder contains initial modelling of any table that will be used in the database.

example :
`User.go`

```go
package models

import (
 "gorm.io/gorm"
)

type User struct {
 gorm.Model
 FullName           string  `gorm:"type:varchar(255)"`
 Email              string  `gorm:"type:varchar(255)"`
 Password           string  `gorm:"type:varchar(255)"`
 }

```

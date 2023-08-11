package repositories

import (
	"fmt"
	"go-restapi-boilerplate/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAllRole(limit, offset int, searchQuery string) (*[]models.MstRole, int64, error)
	FindRoleByID(roleID uint) (*models.MstRole, error)
}

func (r *repository) FindAllRole(limit, offset int, searchQuery string) (*[]models.MstRole, int64, error) {
	var (
		roles     []models.MstRole
		totalRole int64
	)

	trx := r.db.Session(&gorm.Session{})

	if searchQuery != "" {
		trx = trx.Where("role LIKE ?", fmt.Sprintf("%%%s%%", searchQuery))
	}

	trx.Model(&models.MstRole{}).
		Count(&totalRole)

	err := trx.Limit(limit).
		Offset(offset).
		Find(&roles).Error

	return &roles, totalRole, err
}

func (r *repository) FindRoleByID(roleID uint) (*models.MstRole, error) {
	var role models.MstRole
	err := r.db.Where("id = ?", roleID).
		First(&role).Error

	return &role, err
}

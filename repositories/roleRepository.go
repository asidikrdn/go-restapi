package repositories

import (
	"fmt"
	"go-restapi-boilerplate/db/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAllRole(limit, offset int, searchQuery string) (*[]models.MstRole, int64, error)
	FindRoleByID(roleID uint) (*models.MstRole, error)
	CreateRole(role *models.MstRole) (*models.MstRole, error)
	UpdateRole(role *models.MstRole) (*models.MstRole, error)
	DeleteRole(role *models.MstRole) (*models.MstRole, error)
	CheckIsRoleUsed(role *models.MstRole) (bool, error)
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

func (r *repository) CreateRole(role *models.MstRole) (*models.MstRole, error) {
	err := r.db.Create(role).Error

	return role, err
}

func (r *repository) UpdateRole(role *models.MstRole) (*models.MstRole, error) {
	err := r.db.Model(role).Updates(*role).Error

	return role, err
}

func (r *repository) DeleteRole(role *models.MstRole) (*models.MstRole, error) {
	err := r.db.Delete(role).Error

	return role, err
}

func (r *repository) CheckIsRoleUsed(role *models.MstRole) (bool, error) {
	var count int

	err := r.db.Raw("select count(*) from mst_roles mr join mst_users mu on mr.id  = mu.role_id where mr.id = ?", role.ID).Scan(&count).Error

	if count > 0 {
		return true, err
	}

	return false, err
}

package repositories

import (
	"github.com/jinzhu/gorm"
	"lorekeeper/db/models"
)

type RoleRepository struct {
	*Repository
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		Repository: &Repository{Db: db},
	}
}

func (r *RoleRepository) GetByName(name models.RoleName) *models.Role {
	role := &models.Role{}
	r.Db.Where("name = ?", name).First(&role)

	return role
}

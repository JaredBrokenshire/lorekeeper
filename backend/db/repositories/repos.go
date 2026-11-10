package repositories

import (
	"github.com/jinzhu/gorm"
)

type Repos struct {
	DB *gorm.DB

	User *UserRepository
	Role *RoleRepository
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		DB: db,

		User: NewUserRepository(db),
		Role: NewRoleRepository(db),
	}
}

package repositories

import (
	"github.com/jinzhu/gorm"
	"lorekeeper/db/models"
)

type UserRepository struct {
	*Repository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: &Repository{Db: db},
	}
}

func (r *UserRepository) GetById(id interface{}) *models.User {
	user := &models.User{}
	r.Db.
		Preload("Roles").
		Where("id = ?", id).
		First(user)

	return user
}

func (r *UserRepository) GetByUsername(username string) *models.User {
	user := &models.User{}
	r.Db.
		Preload("Roles").
		Where("username = ?", username).
		First(&user)

	return user
}

func (r *UserRepository) GetByEmail(email string) *models.User {
	user := &models.User{}

	r.Db.
		Preload("Roles").
		Where("email = ?", email).
		First(user)

	return user
}

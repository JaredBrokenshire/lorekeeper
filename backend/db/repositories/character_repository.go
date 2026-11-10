package repositories

import (
	"github.com/jinzhu/gorm"
	"lorekeeper/db"
	"lorekeeper/db/models"
)

type CharacterRepository struct {
	*Repository
}

func NewCharacterRepository(db *gorm.DB) *CharacterRepository {
	return &CharacterRepository{
		Repository: &Repository{Db: db},
	}
}

func (r *CharacterRepository) GetCharacters(userID interface{}, page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Character, int, int) {
	var characters []*models.Character

	pageSize, paginateFunc := db.Paginate(page, pageSize)

	r.Db.
		Preload("Race").
		Preload("Classes").
		Preload("Subclasses").
		Scopes(scopes...).
		Scopes(paginateFunc).
		Where("user_id = ?", userID).
		Limit(pageSize).
		Find(&characters)

	return characters, page, pageSize
}

func (r *CharacterRepository) Count(userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) int {
	var count int64

	r.Db.
		Model(&models.Character{}).
		Scopes(scopes...).
		Where("user_id = ?", userID).
		Count(&count)

	return int(count)
}

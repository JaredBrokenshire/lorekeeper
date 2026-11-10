package services

import (
	"github.com/jinzhu/gorm"
	"lorekeeper/db/models"
	"lorekeeper/db/repositories"
)

type CharacterServiceInterface interface {
	List(userID interface{}, page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Character, int, int)
	Count(userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) int
}

type CharacterService struct {
	characterRepo *repositories.CharacterRepository
}

func NewCharacterService(characterRepo *repositories.CharacterRepository) *CharacterService {
	return &CharacterService{characterRepo: characterRepo}
}

func (s *CharacterService) List(userID interface{}, page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Character, int, int) {
	return s.characterRepo.GetCharacters(userID, page, pageSize, scopes)
}

func (s *CharacterService) Count(userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) int {
	return s.characterRepo.Count(userID, scopes)
}

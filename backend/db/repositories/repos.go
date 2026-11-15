package repositories

import (
	"github.com/jinzhu/gorm"
)

type Repos struct {
	DB *gorm.DB

	User          *UserRepository
	Role          *RoleRepository
	Character     *CharacterRepository
	Campaign      *CampaignRepository
	File          *FileRepository
	DocumentChunk *DocumentChunkRepository
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		DB: db,

		User:          NewUserRepository(db),
		Role:          NewRoleRepository(db),
		Character:     NewCharacterRepository(db),
		Campaign:      NewCampaignRepository(db),
		File:          NewFileRepository(db),
		DocumentChunk: NewDocumentChunkRepository(db),
	}
}

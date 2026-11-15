package repositories

import "github.com/jinzhu/gorm"

type DocumentChunkRepository struct {
	*Repository
}

func NewDocumentChunkRepository(db *gorm.DB) *DocumentChunkRepository {
	return &DocumentChunkRepository{
		Repository: &Repository{Db: db},
	}
}

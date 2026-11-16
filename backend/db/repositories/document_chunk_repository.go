package repositories

import (
	"github.com/jinzhu/gorm"
	"lorekeeper/db/models"
)

type DocumentChunkRepository struct {
	*Repository
}

func NewDocumentChunkRepository(db *gorm.DB) *DocumentChunkRepository {
	return &DocumentChunkRepository{
		Repository: &Repository{Db: db},
	}
}

func (r *DocumentChunkRepository) CreateMany(chunks []*models.DocumentChunk) error {
	if len(chunks) == 0 {
		return nil
	}

	tx := r.Db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, chunk := range chunks {
		if chunk == nil {
			continue
		}
		if err := tx.Create(chunk).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *DocumentChunkRepository) GetByCampaignID(campaignID interface{}) ([]*models.DocumentChunk, error) {
	var chunks []*models.DocumentChunk

	err := r.Db.
		Where("campaign_file_id IN (SELECT id FROM campaign_files WHERE campaign_id = ?)", campaignID).
		Order("document_chunks.id ASC").
		Preload("CampaignFile").
		Preload("CampaignFile.File").
		Preload("CampaignFile.File.Author").
		Find(&chunks).Error

	return chunks, err
}

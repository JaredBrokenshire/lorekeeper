package repositories

import (
	"github.com/jinzhu/gorm"
	"lorekeeper/db"
	"lorekeeper/db/models"
)

type CampaignRepository struct {
	*Repository
}

func NewCampaignRepository(db *gorm.DB) *CampaignRepository {
	return &CampaignRepository{
		Repository: &Repository{Db: db},
	}
}

func (r *CampaignRepository) GetCampaigns(userID interface{}, page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Campaign, int, int) {
	var campaigns []*models.Campaign

	pageSize, paginateFunc := db.Paginate(page, pageSize)

	r.Db.
		Joins("LEFT JOIN characters ON characters.campaign_id = campaigns.id").
		Where("dungeon_master_id = ? OR characters.user_id = ?", userID, userID).
		Preload("Characters").
		Preload("Characters.User").
		Preload("Characters.ProfilePicture").
		Preload("DungeonMaster").
		Scopes(scopes...).
		Scopes(paginateFunc).
		Group("campaigns.id").
		Find(&campaigns)

	return campaigns, page, pageSize
}

func (r *CampaignRepository) Count(userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) int {
	var count int64

	r.Db.
		Model(&models.Campaign{}).
		Joins("LEFT JOIN characters ON characters.campaign_id = campaigns.id").
		Where("dungeon_master_id = ? OR characters.user_id = ?", userID, userID).
		Scopes(scopes...).
		Count(&count)

	return int(count)
}

func (r *CampaignRepository) GetByID(id, userID interface{}) *models.Campaign {
	campaign := &models.Campaign{}

	r.Db.
		Joins("LEFT JOIN characters ON characters.campaign_id = campaigns.id").
		Where("campaigns.id = ? AND (dungeon_master_id = ? OR characters.user_id = ?)", id, userID, userID).
		Preload("Files").
		Preload("Files.Author").
		Preload("Characters").
		Preload("Characters.User").
		Preload("Characters.Race").
		Preload("Characters.Classes").
		Preload("Characters.Classes.Logo").
		Preload("Characters.Subclasses").
		Preload("Characters.ProfilePicture").
		Preload("DungeonMaster").
		First(campaign)

	return campaign
}

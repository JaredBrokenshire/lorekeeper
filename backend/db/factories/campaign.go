package factories

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"lorekeeper/db/models"
)

func NewCampaign(db *gorm.DB, campaign *models.Campaign) {
	fillCampaignDefaults(campaign)
	err := db.Create(campaign).Error
	if err != nil {
		panic("error creating campaign in factory: " + err.Error())
	}
}

func fillCampaignDefaults(campaign *models.Campaign) {
	if campaign.Name == "" {
		campaign.Name = random.String(16)
	}
	if campaign.DungeonMasterID == 0 {
		campaign.DungeonMasterID = 1
	}
}

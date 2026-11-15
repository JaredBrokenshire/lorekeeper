package factories

import (
	"github.com/jinzhu/gorm"
	"lorekeeper/db/models"
)

func NewCampaignFile(db *gorm.DB, campaignFile *models.CampaignFile) {
	fillCampaignFileDefaults(campaignFile)
	err := db.Create(campaignFile).Error
	if err != nil {
		panic("error creating campaign file in factory: " + err.Error())
	}
}

func fillCampaignFileDefaults(_ *models.CampaignFile) {
	return
}

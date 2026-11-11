package services

import (
	"github.com/jinzhu/gorm"
	"lorekeeper/db/models"
	"lorekeeper/db/repositories"
)

type CampaignServiceInterface interface {
	List(userID interface{}, page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Campaign, int, int)
	Count(userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) int
}

type CampaignService struct {
	campaignRepo *repositories.CampaignRepository
}

func NewCampaignService(campaignRepo *repositories.CampaignRepository) *CampaignService {
	return &CampaignService{campaignRepo: campaignRepo}
}

func (s *CampaignService) List(userID interface{}, page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Campaign, int, int) {
	return s.campaignRepo.GetCampaigns(userID, page, pageSize, scopes)
}

func (s *CampaignService) Count(userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) int {
	return s.campaignRepo.Count(userID, scopes)
}

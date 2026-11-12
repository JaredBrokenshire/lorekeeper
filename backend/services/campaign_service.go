package services

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
	"lorekeeper/db/models"
	"lorekeeper/db/repositories"
	"lorekeeper/pkg/file_storage"
)

type CampaignServiceInterface interface {
	List(userID uint, page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Campaign, int, int)
	Count(userID uint, scopes []func(db *gorm.DB) *gorm.DB) int
	Get(id interface{}, userID uint) (*models.Campaign, error)
	UploadFile(id interface{}, userID uint, file UploadedFileDTO) (*models.Campaign, error)
}

var (
	ErrCampaignNotFound    = errors.New("campaign not found")
	ErrCampaignInvalidFile = errors.New("invalid file for campaign")
	ErrCampaignFileUpload  = errors.New("upload error for campaign file")
	ErrCampaignFileSave    = errors.New("save error for campaign file")
	ErrCampaignUpdate      = errors.New("update campaign error")
)

type CampaignService struct {
	campaignRepo *repositories.CampaignRepository

	fileService FileServiceInterface
	fileStore   file_storage.Store
}

func NewCampaignService(campaignRepo *repositories.CampaignRepository, fileService FileServiceInterface, fileStore file_storage.Store) *CampaignService {
	return &CampaignService{
		campaignRepo: campaignRepo,
		fileService:  fileService,
		fileStore:    fileStore,
	}
}

func (s *CampaignService) List(userID uint, page, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Campaign, int, int) {
	return s.campaignRepo.GetCampaigns(userID, page, pageSize, scopes)
}

func (s *CampaignService) Count(userID uint, scopes []func(db *gorm.DB) *gorm.DB) int {
	return s.campaignRepo.Count(userID, scopes)
}

func (s *CampaignService) Get(id interface{}, userID uint) (*models.Campaign, error) {
	campaign := s.campaignRepo.GetByID(id, userID)
	if campaign.ID == 0 {
		return nil, ErrCampaignNotFound
	}

	return campaign, nil
}

func (s *CampaignService) UploadFile(id interface{}, userID uint, file UploadedFileDTO) (*models.Campaign, error) {
	campaign, err := s.Get(id, userID)
	if err != nil {
		return nil, ErrCampaignNotFound
	}

	err = s.fileService.Validate(file, []string{"application/pdf"})
	if err != nil {
		return nil, ErrCampaignInvalidFile
	}

	src, err := file.Open()
	if err != nil {
		return nil, ErrCampaignInvalidFile
	}
	defer src.Close()

	name := fmt.Sprintf("%v.pdf", random.String(16))
	path := fmt.Sprintf("/campaigns/%v/session-notes/%v", campaign.ID, name)

	err = s.fileStore.Save(path, src, name)
	if err != nil {
		return nil, ErrCampaignFileUpload
	}

	dto := CreateFileDTO{
		Name:     name,
		Location: path,
		AuthorID: userID,
	}

	fileModel, err := s.fileService.Create(dto)
	if err != nil {
		_ = s.fileStore.Delete(path)
		return nil, ErrCampaignFileSave
	}

	campaign.Files = append(campaign.Files, fileModel)
	err = s.campaignRepo.Update(campaign)
	if err != nil {
		_ = s.fileStore.Delete(path)
		return nil, ErrCampaignUpdate
	}

	return campaign, nil
}

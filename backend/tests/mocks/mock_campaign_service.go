package mocks

import (
	"github.com/jinzhu/gorm"
	"lorekeeper/db/models"
	"lorekeeper/services"
)

type MockCampaignService struct {
	ListCalls []struct {
		UserID   uint
		Page     int
		PageSize int
		Scopes   []func(db *gorm.DB) *gorm.DB
	}

	CountCalls []struct {
		UserID uint
		Scopes []func(db *gorm.DB) *gorm.DB
	}

	GetCalls []struct {
		ID     interface{}
		UserID uint
	}
	GetError error

	UploadFileCalls []struct {
		ID     interface{}
		UserID uint
		File   services.UploadedFileDTO
	}
	UploadFileError error
}

func NewMockCampaignService() *MockCampaignService {
	mock := &MockCampaignService{}
	mock.Reset()
	return mock
}

func (m *MockCampaignService) List(userID uint, page int, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Campaign, int, int) {
	m.ListCalls = append(m.ListCalls, struct {
		UserID   uint
		Page     int
		PageSize int
		Scopes   []func(db *gorm.DB) *gorm.DB
	}{
		userID,
		page,
		pageSize,
		scopes,
	})

	return nil, 0, 0
}

func (m *MockCampaignService) Count(userID uint, scopes []func(db *gorm.DB) *gorm.DB) int {
	m.CountCalls = append(m.CountCalls, struct {
		UserID uint
		Scopes []func(db *gorm.DB) *gorm.DB
	}{
		userID,
		scopes,
	})

	return 0
}

func (m *MockCampaignService) Get(id interface{}, userID uint) (*models.Campaign, error) {
	m.GetCalls = append(m.GetCalls, struct {
		ID     interface{}
		UserID uint
	}{
		id,
		userID,
	})

	return nil, m.GetError
}

func (m *MockCampaignService) UploadFile(id interface{}, userID uint, file services.UploadedFileDTO) (*models.Campaign, error) {
	m.UploadFileCalls = append(m.UploadFileCalls, struct {
		ID     interface{}
		UserID uint
		File   services.UploadedFileDTO
	}{
		id,
		userID,
		file,
	})

	return nil, m.UploadFileError
}

func (m *MockCampaignService) Reset() {
	m.ListCalls = []struct {
		UserID   uint
		Page     int
		PageSize int
		Scopes   []func(db *gorm.DB) *gorm.DB
	}{}

	m.CountCalls = []struct {
		UserID uint
		Scopes []func(db *gorm.DB) *gorm.DB
	}{}

	m.GetCalls = []struct {
		ID     interface{}
		UserID uint
	}{}
	m.GetError = nil

	m.UploadFileCalls = []struct {
		ID     interface{}
		UserID uint
		File   services.UploadedFileDTO
	}{}
	m.UploadFileError = nil
}

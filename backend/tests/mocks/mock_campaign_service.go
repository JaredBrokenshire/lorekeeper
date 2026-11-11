package mocks

import (
	"github.com/jinzhu/gorm"
	"lorekeeper/db/models"
)

type MockCampaignService struct {
	ListCalls []struct {
		UserID   interface{}
		Page     int
		PageSize int
		Scopes   []func(db *gorm.DB) *gorm.DB
	}

	CountCalls []struct {
		UserID interface{}
		Scopes []func(db *gorm.DB) *gorm.DB
	}

	GetCalls []struct {
		ID     interface{}
		UserID interface{}
	}
	GetError error
}

func NewMockCampaignService() *MockCampaignService {
	mock := &MockCampaignService{}
	mock.Reset()
	return mock
}

func (m *MockCampaignService) List(userID interface{}, page int, pageSize int, scopes []func(db *gorm.DB) *gorm.DB) ([]*models.Campaign, int, int) {
	m.ListCalls = append(m.ListCalls, struct {
		UserID   interface{}
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

func (m *MockCampaignService) Count(userID interface{}, scopes []func(db *gorm.DB) *gorm.DB) int {
	m.CountCalls = append(m.CountCalls, struct {
		UserID interface{}
		Scopes []func(db *gorm.DB) *gorm.DB
	}{
		userID,
		scopes,
	})

	return 0
}

func (m *MockCampaignService) Get(id, userID interface{}) (*models.Campaign, error) {
	m.GetCalls = append(m.GetCalls, struct {
		ID     interface{}
		UserID interface{}
	}{
		id,
		userID,
	})

	return nil, m.GetError
}

func (m *MockCampaignService) Reset() {
	m.ListCalls = []struct {
		UserID   interface{}
		Page     int
		PageSize int
		Scopes   []func(db *gorm.DB) *gorm.DB
	}{}

	m.CountCalls = []struct {
		UserID interface{}
		Scopes []func(db *gorm.DB) *gorm.DB
	}{}

	m.GetCalls = []struct {
		ID     interface{}
		UserID interface{}
	}{}
	m.GetError = nil
}

package mocks

import "lorekeeper/services"

type SearchCall struct {
	CampaignID interface{}
	UserID     uint
	Query      string
}

type MockSearchService struct {
	SearchCalls []SearchCall
	SearchError error
}

func NewMockSearchService() *MockSearchService {
	res := &MockSearchService{}
	res.Reset()
	return res
}

func (m *MockSearchService) Search(campaignID interface{}, userID uint, query string) ([]*services.SearchResult, error) {
	m.SearchCalls = append(m.SearchCalls, SearchCall{campaignID, userID, query})
	return nil, m.SearchError
}

func (m *MockSearchService) Reset() {
	m.SearchCalls = []SearchCall{}
	m.SearchError = nil
}

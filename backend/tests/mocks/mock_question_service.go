package mocks

import (
	"context"
	"lorekeeper/services"
)

type AskCall struct {
	CampaignID interface{}
	UserID     uint
	Ctx        context.Context
	Question   string
}

type MockQuestionService struct {
	AskCalls []AskCall
	AskError error
}

func NewMockQuestionService() *MockQuestionService {
	res := &MockQuestionService{}
	res.Reset()
	return res
}

func (m *MockQuestionService) Ask(campaignID interface{}, userID uint, ctx context.Context, question string) (*services.QuestionAnswer, error) {
	m.AskCalls = append(m.AskCalls, AskCall{campaignID, userID, ctx, question})
	return nil, m.AskError

}

func (m *MockQuestionService) Reset() {
	m.AskCalls = []AskCall{}
	m.AskError = nil
}

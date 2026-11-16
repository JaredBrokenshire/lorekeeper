package mocks

import (
	"context"
	"lorekeeper/services"
)

type GenerateCall struct {
	Ctx      context.Context
	Question string
	Context  []services.AnswerContext
}

type MockOllamaService struct {
	GenerateCalls []GenerateCall
	GenerateError error
}

func NewMockOllamaService() *MockOllamaService {
	res := &MockOllamaService{}
	res.Reset()
	return res
}

func (m *MockOllamaService) Generate(ctx context.Context, question string, context []services.AnswerContext) (string, error) {
	m.GenerateCalls = append(m.GenerateCalls, GenerateCall{ctx, question, context})
	return "", m.GenerateError
}

func (m *MockOllamaService) Reset() {
	m.GenerateCalls = []GenerateCall{}
	m.GenerateError = nil
}

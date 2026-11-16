package services

import (
	"context"
)

type QuestionAnswer struct {
	Answer  string
	Sources []*SearchResult
}

type QuestionServiceInterface interface {
	Ask(campaignID interface{}, userID uint, ctx context.Context, question string) (*QuestionAnswer, error)
}

type QuestionService struct {
	searchService SearchServiceInterface
	ollamaService OllamaServiceInterface
}

func NewQuestionService(searchService SearchServiceInterface, ollamaService OllamaServiceInterface) *QuestionService {
	return &QuestionService{searchService: searchService, ollamaService: ollamaService}
}

func (s *QuestionService) Ask(campaignID interface{}, userID uint, ctx context.Context, question string) (*QuestionAnswer, error) {
	searchResults, err := s.searchService.Search(campaignID, userID, question)
	if err != nil {
		return nil, err
	}
	if len(searchResults) == 0 {
		return nil, err
	}

	answerContext := make([]AnswerContext, 0, len(searchResults))
	for _, result := range searchResults {
		if result == nil || result.Chunk == nil {
			continue
		}

		filename := "Unknown Document"

		if result.Chunk.CampaignFile != nil && result.Chunk.CampaignFile.File != nil {
			filename = result.Chunk.CampaignFile.File.Name
		}
		answerContext = append(answerContext, AnswerContext{
			PageNumber: result.Chunk.PageNumber,
			Content:    result.Chunk.Content,
			FileName:   filename,
		})
	}
	if len(answerContext) == 0 {
		return nil, ErrSearchNoResults
	}

	answer, err := s.ollamaService.Generate(ctx, question, answerContext)
	if err != nil {
		return nil, err
	}

	res := &QuestionAnswer{
		Answer:  answer,
		Sources: searchResults,
	}
	return res, nil
}

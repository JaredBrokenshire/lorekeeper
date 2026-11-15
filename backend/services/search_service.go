package services

import (
	"errors"
	"fmt"
	"strings"

	"lorekeeper/db/models"
	"lorekeeper/db/repositories"
	"lorekeeper/pkg/search"
)

const (
	DefaultSearchLimit = 10
	MaxSearchLimit     = 50
	MaxSearchQuerySize = 200
)

var (
	ErrSearchCampaignNotFound = errors.New("campaign not found")
	ErrSearchEmptyQuery       = errors.New("search query is empty")
	ErrSearchQueryTooLong     = errors.New("search query is too long")
	ErrSearchInvalidLimit     = errors.New("search result limit is invalid")
	ErrSearchRepository       = errors.New("error retrieving campaign document chunks")
)

type SearchResult struct {
	Chunk *models.DocumentChunk
	Score float64
}

type SearchServiceInterface interface {
	Search(campaignID interface{}, userID uint, query string, limit int) ([]*SearchResult, error)
}

type SearchService struct {
	campaignRepo *repositories.CampaignRepository
	chunkRepo    *repositories.DocumentChunkRepository
}

func NewSearchService(campaignRepo *repositories.CampaignRepository, chunkRepo *repositories.DocumentChunkRepository) *SearchService {
	return &SearchService{
		campaignRepo: campaignRepo,
		chunkRepo:    chunkRepo,
	}
}

func (s *SearchService) Search(campaignID interface{}, userID uint, query string, limit int) ([]*SearchResult, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, ErrSearchEmptyQuery
	}
	if len(query) > MaxSearchQuerySize {
		return nil, ErrSearchQueryTooLong
	}

	if limit == 0 {
		limit = DefaultSearchLimit
	}
	if limit < 1 || limit > MaxSearchLimit {
		return nil, ErrSearchInvalidLimit
	}

	accessible, err := s.campaignRepo.ExistsForUser(campaignID, userID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrSearchRepository, err)
	}
	if !accessible {
		return nil, ErrSearchCampaignNotFound
	}

	chunks, err := s.chunkRepo.GetByCampaignID(campaignID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrSearchRepository, err)
	}
	if len(chunks) == 0 {
		return []*SearchResult{}, nil
	}

	documents := make([]search.Document, 0, len(chunks))
	chunksByID := make(map[uint]*models.DocumentChunk, len(chunks))
	for _, chunk := range chunks {
		if chunk == nil {
			continue
		}
		documents = append(documents, search.Document{
			ID:   chunk.ID,
			Text: chunk.Content,
		})
		chunksByID[chunk.ID] = chunk
	}

	ranked := search.RankBM25(documents, query, limit)
	results := make([]*SearchResult, 0, len(ranked))
	for _, result := range ranked {
		if chunk, ok := chunksByID[result.ID]; ok {
			results = append(results, &SearchResult{
				Chunk: chunk,
				Score: result.Score,
			})
		}
	}

	return results, nil
}

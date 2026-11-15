package responses

import (
	"lorekeeper/db/models"
	"time"
)

type DocumentChunkResponse struct {
	ID uint `json:"id"`

	Content    string `json:"content"`
	PageNumber int    `json:"page_number"`
	Index      int    `json:"index"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewDocumentChunkResponse(documentChunk *models.DocumentChunk) *DocumentChunkResponse {
	return &DocumentChunkResponse{
		ID: documentChunk.ID,

		Content:    documentChunk.Content,
		PageNumber: documentChunk.PageNumber,
		Index:      documentChunk.Index,

		CreatedAt: documentChunk.CreatedAt,
		UpdatedAt: documentChunk.UpdatedAt,
	}
}

func NewDocumentChunkResponses(documentChunks []*models.DocumentChunk) []DocumentChunkResponse {
	var res []DocumentChunkResponse

	for _, documentChunk := range documentChunks {
		res = append(res, *NewDocumentChunkResponse(documentChunk))
	}

	return res
}

package document

import (
	"github.com/ledongthuc/pdf"
	"strings"
)

type ExtractorInterface interface {
	Extract(reader *pdf.Reader) ([]Page, error)
	Chunk([]Page) ([]Chunk, error)
}

type Extractor struct {
	ChunkSize    int
	ChunkOverlap int
}

func NewExtractor() *Extractor {
	return &Extractor{
		ChunkSize:    50,
		ChunkOverlap: 10,
	}
}

func (e *Extractor) Extract(reader *pdf.Reader) ([]Page, error) {
	totalPages := reader.NumPage()
	pages := make([]Page, 0, totalPages)
	// Cache fonts between pages to avoid repeatedly parsing them.
	fonts := make(map[string]*pdf.Font)

	for pageNumber := 1; pageNumber <= totalPages; pageNumber++ {
		page := reader.Page(pageNumber)

		text, err := page.GetPlainText(fonts)
		if err != nil {
			return nil, err
		}

		pages = append(pages, Page{
			Number: pageNumber,
			Text:   strings.TrimSpace(text),
		})
	}

	return pages, nil
}

func (e *Extractor) Chunk(pages []Page) ([]Chunk, error) {
	var chunks []Chunk
	chunkIndex := 0
	stepSize := e.ChunkSize - e.ChunkOverlap

	for _, page := range pages {
		content := strings.Fields(page.Text)

		if len(content) == 0 {
			continue
		}

		for start := 0; start < len(content); start += stepSize {
			end := start + e.ChunkSize

			if end > len(content) {
				end = len(content)
			}

			chunks = append(chunks, Chunk{
				Index:      chunkIndex,
				Content:    strings.Join(content[start:end], " "),
				PageNumber: page.Number,
			})

			chunkIndex++

			if end == len(content) {
				break
			}
		}
	}

	return chunks, nil
}

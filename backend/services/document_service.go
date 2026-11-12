package services

import (
	"errors"
	"github.com/ledongthuc/pdf"
	"strings"
)

type DocumentServiceInterface interface {
	Extract(path string) ([]Page, error)
}

var (
	ErrDocumentEmptyPath = errors.New("document path is empty")
	ErrDocumentOpen      = errors.New("error opening document")
	ErrDocumentPageRead  = errors.New("error reading page")
)

type DocumentService struct{}

func NewDocumentService() *DocumentService {
	return &DocumentService{}
}

type Page struct {
	Number int
	Text   string
}

func (s *DocumentService) Extract(path string) ([]Page, error) {
	if strings.TrimSpace(path) == "" {
		return nil, ErrDocumentEmptyPath
	}

	file, reader, err := pdf.Open(path)
	if err != nil {
		return nil, ErrDocumentOpen
	}
	defer file.Close()

	totalPages := reader.NumPage()

	pages := make([]Page, 0, totalPages)
	// Cache fonts between pages to avoid repeatedly parsing them.
	fonts := make(map[string]*pdf.Font)

	for pageNumber := 1; pageNumber <= totalPages; pageNumber++ {
		page := reader.Page(pageNumber)

		text, err := page.GetPlainText(fonts)
		if err != nil {
			return nil, ErrDocumentPageRead
		}

		pages = append(pages, Page{
			Number: pageNumber,
			Text:   strings.TrimSpace(text),
		})
	}

	return pages, nil
}

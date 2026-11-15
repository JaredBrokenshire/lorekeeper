package dependency_tests

import (
	"fmt"
	"github.com/ledongthuc/pdf"
	"github.com/stretchr/testify/assert"
	"lorekeeper/pkg/document"
	"path/filepath"
	"strings"
	"testing"
)

type TextPart struct {
	Page int
	Text string
}

type Expected struct {
	Pages     int
	TextParts []TextPart
	Error     error
}

func TestDocument_Extract(t *testing.T) {
	extractor := document.NewExtractor()

	cases := []struct {
		Name     string
		Path     string
		Expected Expected
	}{
		{
			Name: "Can extract single page PDF",
			Path: "test.pdf",
			Expected: Expected{
				Pages: 1,
				TextParts: []TextPart{
					{
						Page: 1,
						Text: "Brackenford",
					},
					{
						Page: 1,
						Text: "The drowned creatures",
					},
					{
						Page: 1,
						Text: "Recovered token bearing three waves.",
					},
				},
			},
		},
		{
			Name: "Can extract multi-page PDF",
			Path: "test-multi-page.pdf",
			Expected: Expected{
				Pages: 2,
				TextParts: []TextPart{
					{
						Page: 1,
						Text: "Stayed at the Tallow & Thorn Inn",
					},
					{
						Page: 1,
						Text: " Nobody admits to ringing it.",
					},
					{
						Page: 1,
						Text: "Scratches on the inside of the bell housing.",
					},
					{
						Page: 2,
						Text: "Several dead figures emerged from the cellar beneath the tower",
					},
					{
						Page: 2,
						Text: "5. Drowned 2",
					},
					{
						Page: 2,
						Text: "Why are the bell-ringers disappearing?",
					},
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			_, r, _ := pdf.Open(filepath.Join("../../assets", test.Path))

			pages, err := extractor.Extract(r)
			assert.Equal(t, test.Expected.Error, err)

			assert.Equal(t, test.Expected.Pages, len(pages))

			for _, textPart := range test.Expected.TextParts {
				assert.Contains(t, pages[textPart.Page-1].Text, textPart.Text)
			}
		})
	}
}

func TestDocument_Chunk(t *testing.T) {
	extractor := document.NewExtractor()

	words := make([]string, 500)
	for i := range words {
		words[i] = fmt.Sprintf("word%d", i+1)
	}

	pages := []document.Page{
		{
			Number: 1,
			Text:   strings.Join(words, " "),
		},
	}

	chunks, err := extractor.Chunk(pages)
	assert.NoError(t, err)

	wordCount := len(words)
	expectedChunks := (wordCount + (extractor.ChunkSize - extractor.ChunkOverlap) - 1) / (extractor.ChunkSize - extractor.ChunkOverlap)

	assert.Len(t, chunks, expectedChunks)
	assert.Equal(t, "word1", strings.Fields(chunks[0].Content)[0])
	assert.Equal(t, fmt.Sprintf("word%v", extractor.ChunkSize), strings.Fields(chunks[0].Content)[extractor.ChunkSize-1])
	assert.Equal(t, fmt.Sprintf("word%v", extractor.ChunkSize-extractor.ChunkOverlap+1), strings.Fields(chunks[1].Content)[0])
	assert.Equal(t, fmt.Sprintf("word%v", 2*extractor.ChunkSize-extractor.ChunkOverlap), strings.Fields(chunks[1].Content)[extractor.ChunkSize-1])
	lastChunk := chunks[len(chunks)-1].Content
	lastWord := strings.Fields(lastChunk)[len(strings.Fields(lastChunk))-1]
	assert.Equal(t, "word500", lastWord)
}

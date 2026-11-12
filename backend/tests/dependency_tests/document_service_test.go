package dependency_tests

import (
	"github.com/stretchr/testify/assert"
	"lorekeeper/services"
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

func TestDocumentService_Extract(t *testing.T) {
	documentService := services.NewDocumentService()

	cases := []struct {
		Name     string
		Path     string
		Expected Expected
	}{
		{
			Name: "Can't extract PDF with no path",
			Path: "",
			Expected: Expected{
				Error: services.ErrDocumentEmptyPath,
			},
		},
		{
			Name: "Can't extract PDF that doesn't exist",
			Path: "test-file-does-not-exist.pdf",
			Expected: Expected{
				Error: services.ErrDocumentOpen,
			},
		},
		{
			Name: "Can't extract empty PDF",
			Path: "../../assets/test-empty.pdf",
			Expected: Expected{
				Error: services.ErrDocumentOpen,
			},
		},
		{
			Name: "Can't extract corrupted PDF",
			Path: "../../assets/test-corrupted.pdf",
			Expected: Expected{
				Error: services.ErrDocumentOpen,
			},
		},
		{
			Name: "Can extract single page PDF",
			Path: "../../assets/test.pdf",
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
			Path: "../../assets/test-multi-page.pdf",
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
			pages, err := documentService.Extract(test.Path)
			assert.Equal(t, test.Expected.Error, err)

			assert.Equal(t, test.Expected.Pages, len(pages))

			for _, textPart := range test.Expected.TextParts {
				assert.Contains(t, pages[textPart.Page-1].Text, textPart.Text)
			}
		})
	}
}

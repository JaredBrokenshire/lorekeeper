package dependency_tests

import (
	"lorekeeper/pkg/search"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRankBM25_RanksRelevantDocuments(t *testing.T) {
	documents := []search.Document{
		{ID: 1, Text: "The party entered the old tower."},
		{ID: 2, Text: "The drowned creatures emerged beneath the tower and attacked the party."},
		{ID: 3, Text: "A merchant sold the party a map."},
	}

	results := search.RankBM25(documents, "drowned creatures tower")

	if assert.GreaterOrEqual(t, len(results), 1) {
		assert.Equal(t, uint(2), results[0].ID)
		assert.Greater(t, results[0].Score, 0.0)
	}
}

func TestRankBM25_EmptyQueryOrNoMatchesReturnsEmpty(t *testing.T) {
	documents := []search.Document{{ID: 1, Text: "A quiet village"}}

	assert.Empty(t, search.RankBM25(documents, "  "))
	assert.Empty(t, search.RankBM25(documents, "spaceship"))
	assert.Empty(t, search.RankBM25(documents, "village"))
}

func TestRankBM25_TokenizesUnicodeAndPunctuation(t *testing.T) {
	documents := []search.Document{
		{ID: 1, Text: "Brackenford's drowned-creatures!"},
	}

	results := search.RankBM25(documents, "BRACKENFORD drowned")

	if assert.Len(t, results, 1) {
		assert.Equal(t, uint(1), results[0].ID)
	}
}

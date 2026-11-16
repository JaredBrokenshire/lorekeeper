package search

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

var (
	MaxSearchLimit = 10
	tokenPattern   = regexp.MustCompile(`[\p{L}\p{N}]+`)
)

type Document struct {
	ID   uint
	Text string
}

type Result struct {
	ID    uint
	Score float64
}

func RankBM25(documents []Document, query string) []Result {
	queryTerms := uniqueTerms(tokenize(query))
	if len(documents) == 0 || len(queryTerms) == 0 {
		return []Result{}
	}

	type indexedDocument struct {
		id     uint
		terms  map[string]int
		length int
	}

	indexed := make([]indexedDocument, 0, len(documents))
	documentFrequency := make(map[string]int)
	totalLength := 0

	for _, doc := range documents {
		tokens := tokenize(doc.Text)
		termFrequency := make(map[string]int, len(tokens))
		for _, token := range tokens {
			termFrequency[token]++
		}

		indexed = append(indexed, indexedDocument{
			id:     doc.ID,
			terms:  termFrequency,
			length: len(tokens),
		})
		totalLength += len(tokens)

		for term := range termFrequency {
			documentFrequency[term]++
		}
	}

	if len(indexed) == 0 {
		return []Result{}
	}

	averageLength := float64(totalLength) / float64(len(indexed))
	if averageLength == 0 {
		return []Result{}
	}

	const (
		k1 = 1.5
		b  = 0.75
	)

	results := make([]Result, 0, len(indexed))
	documentCount := float64(len(indexed))

	for _, doc := range indexed {
		score := 0.0

		for _, term := range queryTerms {
			tf := float64(doc.terms[term])
			if tf == 0 {
				continue
			}

			df := float64(documentFrequency[term])
			idf := math.Log(1 + (documentCount-df+0.5)/(df+0.5))
			lengthNormalisation := k1 * (1 - b + b*float64(doc.length)/averageLength)
			termScore := idf * (tf * (k1 + 1)) / (tf + lengthNormalisation)
			score += termScore
		}

		if score > 0.25 {
			results = append(results, Result{ID: doc.id, Score: score})
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if len(results) > MaxSearchLimit {
		results = results[:MaxSearchLimit]
	}

	return results
}

func tokenize(value string) []string {
	matches := tokenPattern.FindAllString(strings.ToLower(value), -1)
	return matches
}

func uniqueTerms(terms []string) []string {
	seen := make(map[string]struct{}, len(terms))
	unique := make([]string, 0, len(terms))
	for _, term := range terms {
		if _, exists := seen[term]; exists {
			continue
		}
		seen[term] = struct{}{}
		unique = append(unique, term)
	}
	return unique
}

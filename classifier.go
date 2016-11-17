package classifier

import (
	"regexp"
	"strings"
)

// Classifier provides a simple interface for different text classifiers
type Classifier interface {
	// Train allows clients to train the classifier
	Train(doc string, category string) error
	// Classify performs a classification on the input corpus and assumes that
	// the underlying classifier has been trained.
	Classify(doc string) (string, error)
}

// IsAWord is a predicate to determine whether or not a token should be considered a
// valid word (i.e. greater than two but less than 20 characters).
func IsAWord(v string) bool {
	return len(v) > 2 && len(v) < 20
}

// Tokenize extracts and normalizes all words from a text corpus
func Tokenize(doc string) ([]string, error) {
	tokenizer, err := regexp.Compile("\\W+")

	if err != nil {
		return nil, err
	}

	tokens := tokenizer.Split(doc, -1)
	return Map(Filter(tokens, IsAWord), strings.ToLower), nil
}

// WordCounts extracts term frequencies from a text corpus
func WordCounts(doc string) (map[string]int, error) {
	tokens, err := Tokenize(doc)

	if err != nil {
		return nil, err
	}

	wc := make(map[string]int)
	for _, token := range tokens {
		wc[token] = wc[token] + 1
	}
	return wc, nil
}

// Map applies f to each element of the supplied input slice
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Filter removes elements from the input slice where the supplied predicate
// is satisfied
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

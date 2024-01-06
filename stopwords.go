package classifier

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

var (
	stopwords = map[string]struct{}{
		"a": {}, "able": {}, "about": {}, "across": {}, "after": {}, "all": {}, "almost": {}, "also": {}, "am": {}, "among": {}, "an": {}, "and": {}, "any": {}, "are": {}, "as": {}, "at": {},
		"be": {}, "because": {}, "been": {}, "but": {}, "by": {}, "can": {}, "cannot": {}, "could": {}, "dear": {}, "did": {}, "do": {}, "does": {}, "either": {}, "else": {}, "ever": {},
		"every": {}, "for": {}, "from": {}, "get": {}, "got": {}, "had": {}, "has": {}, "have": {}, "he": {}, "her": {}, "hers": {}, "him": {}, "his": {}, "how": {}, "however": {}, "i": {},
		"if": {}, "in": {}, "into": {}, "is": {}, "it": {}, "its": {}, "just": {}, "least": {}, "let": {}, "like": {}, "likely": {}, "may": {}, "me": {}, "might": {}, "most": {}, "must": {},
		"my": {}, "neither": {}, "no": {}, "nor": {}, "not": {}, "of": {}, "off": {}, "often": {}, "on": {}, "only": {}, "or": {}, "other": {}, "our": {}, "own": {}, "rather": {}, "said": {},
		"say": {}, "says": {}, "she": {}, "should": {}, "since": {}, "so": {}, "some": {}, "than": {}, "that": {}, "the": {}, "their": {}, "them": {}, "then": {}, "there": {}, "these": {},
		"they": {}, "this": {}, "tis": {}, "to": {}, "too": {}, "twas": {}, "us": {}, "wants": {}, "was": {}, "we": {}, "were": {}, "what": {}, "when": {}, "where": {}, "which": {}, "while": {},
		"who": {}, "whom": {}, "why": {}, "will": {}, "with": {}, "would": {}, "yet": {}, "you": {}, "your": {},
	}
)

// IsStopWord checks against a list of known english stop words and returns true if v is a
// stop word; false otherwise
func IsStopWord(v string) bool {
	if _, ok := stopwords[strings.ToLower(v)]; ok {
		return true
	}
	return false
}

// IsNotStopWord is the inverse function of IsStopWord
func IsNotStopWord(v string) bool {
	return !IsStopWord(v)
}

func LoadStopWords(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		stopwords[record[0]] = struct{}{}
	}
	return nil
}
package classifier

import "testing"

func TestStopWords(t *testing.T) {
	t.Run("Stopword", func(t *testing.T) {
		sample := []string{"a", "is", "the"}
		for _, v := range sample {
			if IsNotStopWord(v) {
				t.Errorf("%s was not identified as a stop word", v)
			}
		}
	})
	t.Run("Other", func(t *testing.T) {
		sample := []string{"hello", "world"}
		for _, v := range sample {
			if IsStopWord(v) {
				t.Errorf("%s was incorrectly identified as a stop word", v)
			}
		}
	})
}

func TestLoadStopWords(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"load data from file",
			args{"data/stop_words.csv"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadStopWords(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("LoadStopWords() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !IsStopWord("что") {
				t.Error("IsStopWord() error")
			}
		})
	}
}

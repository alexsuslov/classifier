// (c) 2023 Alex Suslov
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package naive

import (
	"testing"
)

func TestClassifier_Save(t *testing.T) {
	c := New()
	c.TrainString("The quick brown fox jumped over the lazy dog", "ham")
	c.TrainString("Earn a degree online", "ham")
	c.TrainString("Earn cash quick online", "spam")
	type fields struct {
		C *Classifier
	}
	type args struct {
		filename  string
		filename1 string
	}
	tests := []struct {
		name               string
		C                  *Classifier
		args               args
		wantClassification string
		wantErr            bool
	}{
		{"SaveFeat2cat", c,
			args{
				"../data/test.fea",
				"../data/test.cat",
			},
			"spam",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.C
			if err := c.SaveFeat2cat(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("SaveFeat2cat() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := c.SaveCatCount(tt.args.filename1); (err != nil) != tt.wantErr {
				t.Errorf("SaveFeat2cat() error = %v, wantErr %v", err, tt.wantErr)
			}

			classification, err := c.ClassifyString("Earn your masters degree online")
			if err != nil {
				t.Errorf("ClassifyString() error = %v, wantErr %v", err, tt.wantErr)

			}
			if classification != tt.wantClassification {
				t.Errorf("Classification error = %v, want %v", classification, tt.wantClassification)
			}

		})
	}
}

func TestClassifier_Load(t *testing.T) {
	c := New()

	type fields struct {
		C *Classifier
	}
	type args struct {
		filename  string
		filename1 string
	}
	tests := []struct {
		name               string
		C                  *Classifier
		args               args
		wantClassification string
		wantErr            bool
		wantErr1           bool
		wantErr2           bool
	}{
		{"SaveFeat2cat", c,
			args{
				"../data/test.fea",
				"../data/test.cat",
			},
			"spam",
			false, false, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.C
			if err := c.LoadFeat2cat(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("SaveFeat2cat() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := c.LoadCatCount(tt.args.filename1); (err != nil) != tt.wantErr1 {
				t.Errorf("SaveFeat2cat() error = %v, wantErr %v", err, tt.wantErr)
			}
			classification, err := c.ClassifyString("Earn your masters degree online")
			if err != nil {
				t.Errorf("ClassifyString() error = %v, wantErr %v", err, tt.wantErr1)

			}
			if classification != tt.wantClassification {
				t.Errorf("Classification error = %v, want %v", classification, tt.wantClassification)
			}

		})
	}
}

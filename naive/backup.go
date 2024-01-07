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
	"encoding/gob"
	"os"
)

func (c *Classifier) SaveFeat2cat(filename string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	w, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	return gob.NewEncoder(w).Encode(c.feat2cat)
}

func (c *Classifier) SaveCatCount(filename string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	w, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	return gob.NewEncoder(w).Encode(c.catCount)
}

func (c *Classifier) LoadFeat2cat(filename string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	return gob.NewDecoder(r).Decode(&c.feat2cat)
}

func (c *Classifier) LoadCatCount(filename string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	return gob.NewDecoder(r).Decode(&c.catCount)
}

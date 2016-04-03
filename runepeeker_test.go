// MIT License
//
// Copyright (c) 2016 Jay Hirata
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
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

package runepeeker_test

import (
	"github.com/jhizzle/runepeeker"
	"io"
	"strings"
	"testing"
)

func TestRunePeekAndRead(t *testing.T) {
	samples := []struct {
		str    string
		read   int
		peek   int
		expect string
		err    error
	}{
		{
			str:    "Héllô ¥º®£Ð",
			peek:   1,
			expect: "H",
		},
		{
			read:   1,
			expect: "H",
		},
		{
			read:   1,
			expect: "é",
		},
		{
			read:   1,
			expect: "l",
		},
		{
			peek:   4,
			expect: "lô ¥",
		},
		{
			read:   4,
			expect: "lô ¥",
		},
		{
			read:   10,
			expect: "º®£Ð",
			err:    io.EOF,
		},
		{
			peek:   10,
			expect: "",
			err:    io.EOF,
		},
		{
			str:    "Héllô ¥º®£Ð",
			peek:   30,
			expect: "Héllô ¥º®£Ð",
			err:    io.EOF,
		},
		{
			peek:   5,
			expect: "Héllô",
		},
		{
			peek:   11,
			expect: "Héllô ¥º®£Ð",
		},
		{
			peek:   12,
			expect: "Héllô ¥º®£Ð",
			err:    io.EOF,
		},
		{
			read:   11,
			expect: "Héllô ¥º®£Ð",
		},
		{
			read:   1,
			expect: "",
			err:    io.EOF,
		},
		{
			str:    "Héllô ¥º®£Ð",
			read:   10,
			expect: "Héllô ¥º®£",
		},
		{
			read:   1,
			expect: "Ð",
		},
		{
			peek:   1,
			expect: "",
			err:    io.EOF,
		},
		{
			read:   1,
			expect: "",
			err:    io.EOF,
		},
	}
	rp := runepeeker.New(strings.NewReader(""))

	var actual string
	var actualErr error
	for n, test := range samples {
		if test.str != "" {
			rp = runepeeker.New(strings.NewReader(test.str))
		}
		if test.read != 0 {
			// handle reads
			actual = ""
			for i := 0; i < test.read; i++ {
				c, _, err := rp.ReadRune()
				actualErr = err
				if err != nil {
					break
				}
				actual = actual + string(c)
			}
		} else {
			// handle peeks
			c, err := rp.PeekRune(test.peek)
			actual, actualErr = string(c), err
		}

		if string(actual) != test.expect {
			t.Errorf("%3d: Unexpected rune,  Expected: %s (%#q), Actual: %s (%#q)\n",
				n,
				test.expect,
				test.expect,
				actual,
				actual,
			)
		}
		if actualErr != test.err {
			t.Errorf("%3d: Unexpected error, Expected: %s (%#q), Actual: %s (%#q)\n",
				n,
				test.err,
				test.err,
				actualErr,
				actualErr,
			)
		}
	}
}

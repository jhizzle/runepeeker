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
//
// RunePeeker - Peek and read runes from a io.RuneReader
// Copyright (C) 2016  Jay Hirata

// Package runepeeker provides a wrapper around a RuneReader that allows you to
// peek at the runes availabile.
package runepeeker

import (
	"io"
	"unicode/utf8"
)

// RunePeeker lets you read and peek runes.
type RunePeeker struct {
	rr  io.RuneReader
	buf []rune
}

// New takes an io.RuneReader and wraps it to provide a RunePeeker.
func New(reader io.RuneReader) *RunePeeker {
	return &RunePeeker{reader, nil}
}

// ReadRune reads a rune. It returns the rune, the size of the rune (in
// bytes), and error if the read was not successful.
func (r *RunePeeker) ReadRune() (rune, int, error) {
	var c rune
	if len(r.buf) > 0 {
		c, r.buf = r.buf[0], r.buf[1:]
		return c, utf8.RuneLen(c), nil
	}
	return r.rr.ReadRune()
}

// PeekRune will peek at the next n runes. It returns a slice of runes and an
// error. If the slice is less than n in length, than the error is non nil.
func (r *RunePeeker) PeekRune(n int) ([]rune, error) {
	var returnbuf []rune
	var peekbuf []rune
	var c rune
	var err error = nil

	for i := 0; i < n; i++ {

		if len(r.buf) > 0 {
			c, r.buf, err = r.buf[0], r.buf[1:], nil
		} else {
			c, _, err = r.rr.ReadRune()
		}
		if err != nil {
			break
		}
		returnbuf = append(returnbuf, c)
		peekbuf = append(peekbuf, c)
	}

	r.buf = append(returnbuf, r.buf...)

	return peekbuf, err
}

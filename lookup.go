package wig

import (
	"io"
)

type LookupReader struct {
	r     io.Reader
	table map[byte]byte
}

func NewLookupReader(r io.Reader) LookupReader {
	t := make(map[byte]byte, 52)
	for _, c := range []int{65, 97} {
		for i := 0; i < 26; i++ {
			x := byte(i + c)
			y := byte((i+13)%26 + c)
			t[x] = y
		}
	}
	return LookupReader{
		r:     r,
		table: t,
	}
}

func (r LookupReader) Lookup(b byte) byte {
	i, ok := r.table[b]
	if ok {
		return i
	} else {
		return b
	}
	return b
}

func (r LookupReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		p[i] = r.Lookup(p[i])
	}
	return n, err
}

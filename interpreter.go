package wig

import (
	"io"
)

type InterpretReader struct {
	r io.Reader
}

func NewInterpretReader(r io.Reader) InterpretReader {
	return InterpretReader{r}
}

func (r InterpretReader) Interpret(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

func (r InterpretReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		p[i] = r.Interpret(p[i])
	}
	return n, err
}

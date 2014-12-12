package wig

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func ExampleLookup() {
	s := strings.NewReader("Jevgr va Tb!\nJevgr va Tb!")
	r := NewLookupReader(s)
	io.Copy(os.Stdout, &r)
	// Output:
	// Write in Go!
	// Write in Go!
}

func BenchmarkLookup(b *testing.B) {
	s := strings.NewReader("Jevgr va Tb!\nJevgr va Tb!")
	for i := 0; i < b.N; i++ {
		r := NewLookupReader(s)
		buf := make([]byte, s.Len())
		r.Read(buf)
	}
}

func lookupLongString(s string, n int) string {
	b := bytes.NewBuffer(nil)
	w := bufio.NewWriterSize(b, len(s)*n)
	for i := 0; i < n; i++ {
		w.WriteString(s)
	}
	w.Flush()
	return b.String()
}

func BenchmarkLookupLongString(b *testing.B) {
	b.StopTimer()
	s := strings.NewReader(lookupLongString("Jevgr va Tb!\n", 1000000))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r := NewLookupReader(s)
		buf := make([]byte, s.Len())
		r.Read(buf)
	}
}

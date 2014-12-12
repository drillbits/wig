package wig

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func ExampleInterpreter() {
	s := strings.NewReader("Jevgr va Tb!\nJevgr va Tb!")
	r := NewInterpretReader(s)
	io.Copy(os.Stdout, &r)
	// Output:
	// Write in Go!
	// Write in Go!
}

func BenchmarkInterpreter(b *testing.B) {
	s := strings.NewReader("Jevgr va Tb!\nJevgr va Tb!")
	for i := 0; i < b.N; i++ {
		r := NewInterpretReader(s)
		buf := make([]byte, s.Len())
		r.Read(buf)
	}
}

func InterpreterLongString(s string, n int) string {
	b := bytes.NewBuffer(nil)
	w := bufio.NewWriterSize(b, len(s)*n)
	for i := 0; i < n; i++ {
		w.WriteString(s)
	}
	w.Flush()
	return b.String()
}

func BenchmarkInterpreterLongString(b *testing.B) {
	b.StopTimer()
	s := strings.NewReader(InterpreterLongString("Jevgr va Tb!\n", 1000000))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r := NewInterpretReader(s)
		buf := make([]byte, s.Len())
		r.Read(buf)
	}
}

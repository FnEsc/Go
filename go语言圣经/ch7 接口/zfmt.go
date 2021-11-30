package main

import (
	"bytes"
	"io"
	"os"
)

func Fprintf(w io.Writer, format string, args ...interface{}) (int, error) {
	return 0, nil
}

func Printf(format string, args ...interface{}) (int, error) {
	return Fprintf(os.Stdout, format, args...)
}

func Sprintf(format string, args ...interface{}) string {
	var buf bytes.Buffer
	Fprintf(&buf, format, args...)
	return buf.String()
}

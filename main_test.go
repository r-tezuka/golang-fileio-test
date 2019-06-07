package main

import (
	"bytes"
	"testing"
)

func TestAddText(t *testing.T) {
	r := bytes.NewBufferString("foo\n")
	w := new(bytes.Buffer)

	AddText(r, w)

	expected := []byte("foo\nhoge")

	if bytes.Compare(expected, w.Bytes()) != 0 {
		t.Fatal("not matched")
	}
}

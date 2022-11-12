package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4
	actual := count(b, false, false)

	if actual != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, actual)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

	exp := 3
	actual := count(b, true, false)

	if actual != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, actual)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

	exp := 35
	actual := count(b, false, true)

	if actual != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, actual)
	}
}

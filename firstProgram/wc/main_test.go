package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	countLines := false
	countBytes := false
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4        // expected
	res := count(b, countLines, countBytes) // actual (result)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

// TestCountLines tests the count function set to count lines
func TestCountLines(t *testing.T) {
	countLines := true
	countBytes := false
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := 3
	res := count(b, countLines, countBytes)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

// TestCountBytes tests the count function set to count bytes
func TestCountBytes(t *testing.T) {
	countLines := false
	countBytes := true
	b := bytes.NewBufferString("word1 word2\n")
	exp := 12
	res := count(b, countLines, countBytes)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
package main

import "testing"

func TestHexConversion(t *testing.T) {
	input := "1E (hex) files"
	expected := "30 files"
	result := ProcessText(input)
	if result != expected {
		t.Errorf("got %s, want %s", result, expected)
	}
}

func TestBinConversion(t *testing.T) {
	input := "10 (bin) years"
	expected := "2 years"
	result := ProcessText(input)
	if result != expected {
		t.Errorf("got %s, want %s", result, expected)
	}
}

func TestUppercase(t *testing.T) {
	input := "go (up)"
	expected := "GO"
	result := ProcessText(input)
	if result != expected {
		t.Errorf("got %s, want %s", result, expected)
	}
}

func TestLowercaseMultiple(t *testing.T) {
	input := "HELLO WORLD (low, 2)"
	expected := "hello world"
	result := ProcessText(input)
	if result != expected {
		t.Errorf("got %s, want %s", result, expected)
	}
}

func TestCapitalize(t *testing.T) {
	input := "bridge (cap)"
	expected := "Bridge"
	result := ProcessText(input)
	if result != expected {
		t.Errorf("got %s, want %s", result, expected)
	}
}

func TestArticles(t *testing.T) {
	input := "a amazing story"
	expected := "an amazing story"
	result := ProcessText(input)
	if result != expected {
		t.Errorf("got %s, want %s", result, expected)
	}
}

func TestPunctuation(t *testing.T) {
	input := "Hello , world !"
	expected := "Hello, world!"
	result := ProcessText(input)
	if result != expected {
		t.Errorf("got %s, want %s", result, expected)
	}
}

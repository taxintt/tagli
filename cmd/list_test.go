package cmd

import (
	"testing"
)

func TestPrintPlainFormat(t *testing.T) {
	plain := printPlainFormat(iter)
	expected := `v1.0
v2.0
`
	if plain != expected {
		t.Errorf("Expected %s, got %s", expected, plain)
	}
}

func TestPrintJsonFormat(t *testing.T) {
	tags := make(map[string]string)
	tags["v1.0"] = "1234567890"
	tags["v2.0"] = "0987654321"

	json := printJsonFormat(iter)
	expected := `{
		  "v1.0": "1234567890",
		  "v2.0": "0987654321"
		}`
	if json != expected {
		t.Errorf("Expected %s, got %s", expected, json)
	}
}

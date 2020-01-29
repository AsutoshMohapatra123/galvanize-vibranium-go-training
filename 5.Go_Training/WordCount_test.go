package main

import (
	"testing"
)

func TestwordCount(t *testing.T) {

	var input = "Australia Canada Germany Australia Japan Canada"

	output := make(map[string]int)
	output = wordCount(input)

	if output["Canada"] != 2 {
		t.Errorf("Canada Word count Mismatch")
	} else if output["Germany"] != 1 {
		t.Errorf("Germany Word count Mismatch")
	} else if output["Japan"] != 1 {
		t.Errorf("Japan Word count Mismatch")
	} else if output["Australia"] != 1 {
		t.Errorf("Australia Word count Mismatch")
	}

}

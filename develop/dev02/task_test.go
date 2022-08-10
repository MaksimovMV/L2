package main

import "testing"

func TestUnpackString(t *testing.T) {
	data := []struct {
		input  string
		output string
	}{
		{
			input:  "a4bc2d5e",
			output: "aaaabccddddde",
		},
		{
			input:  "abcd",
			output: "abcd",
		},
		{
			input:  "45",
			output: "",
		},
		{
			input:  "",
			output: "",
		},
	}

	for i, testCase := range data {
		res := unpackString(testCase.input)
		if res != testCase.output {
			t.Errorf("test %v, Expected %v, got %v", i, testCase.output, res)
		}
	}
}

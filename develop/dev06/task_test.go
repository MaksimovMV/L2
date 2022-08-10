package main

import (
	"fmt"
	"testing"
)

func TestCut(t *testing.T) {
	data := []struct {
		a      string
		f      int
		d      string
		s      bool
		err    error
		result string
	}{
		{
			a:      "дек\t6\tAndroid\t4",
			f:      1,
			d:      "\t",
			s:      false,
			result: "дек",
		},
		{
			a:      "дек\t6\tAndroid\t4",
			f:      3,
			d:      "\t",
			s:      false,
			result: "Android",
		},
		{
			a:      "дек 6 Android\t4",
			f:      3,
			d:      " ",
			s:      false,
			result: "Android\t4",
		},
		{
			a:      "дек 6 Android 4",
			f:      1,
			d:      "\t",
			s:      true,
			result: "",
		},
		{
			a:      "дек\t6\tAndroid\t4",
			f:      5,
			d:      "\t",
			s:      false,
			err:    fmt.Errorf("некорректное поле f"),
			result: "",
		},
	}

	for i, testCase := range data {
		res, err := cut(testCase.a, testCase.f, testCase.d, testCase.s)
		if res != testCase.result {
			t.Errorf("test %v, Expected %v, got %v", i, testCase.result, res)
		}
		if err != nil && err.Error() != testCase.err.Error() {
			t.Errorf("Expected %+v, got %+v", testCase.err.Error(), err.Error())
		}
	}
}

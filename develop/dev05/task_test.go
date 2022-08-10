package main

import (
	"reflect"
	"testing"
)

func TestFindNumsOfLines(t *testing.T) {
	testArr := []string{
		"янв 15 Downloads 5b",
		"дек 6 Android 4",
		"июн 10 Sources 45",
		"окт 31 VirtualBox 90",
		"окт 31 VirtualBox 90",
		"янв 13 Lightworks 9k",
		"янв 11 Pictures 12",
		"янв 11 P|pictures 12"}

	data := []struct {
		arr        []string
		pattern    string
		after      int
		before     int
		context    int
		count      bool
		ignoreCase bool
		invert     bool
		fixed      bool
		lineNum    bool
		result     []int
	}{
		{
			arr:        testArr,
			pattern:    "VirtualBox",
			after:      2,
			before:     0,
			context:    0,
			ignoreCase: false,
			invert:     false,
			fixed:      false,
			result:     []int{3, 4, 5, 6},
		},
		{
			arr:        testArr,
			pattern:    "VirtualBox",
			after:      0,
			before:     2,
			context:    0,
			ignoreCase: false,
			invert:     false,
			fixed:      false,
			result:     []int{1, 2, 3, 4},
		},
		{
			arr:        testArr,
			pattern:    "VirtualBox",
			after:      0,
			before:     0,
			context:    1,
			ignoreCase: false,
			invert:     false,
			fixed:      false,
			result:     []int{2, 3, 4, 5},
		},
		{
			arr:        testArr,
			pattern:    "VirtualBox",
			after:      1,
			before:     1,
			context:    4,
			ignoreCase: false,
			invert:     false,
			fixed:      false,
			result:     []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			arr:        testArr,
			pattern:    "virtualbox",
			after:      0,
			before:     0,
			context:    0,
			ignoreCase: true,
			invert:     false,
			fixed:      false,
			result:     []int{3, 4},
		},
		{
			arr:        testArr,
			pattern:    "VirtualBox",
			after:      0,
			before:     0,
			context:    0,
			ignoreCase: false,
			invert:     true,
			fixed:      false,
			result:     []int{0, 1, 2, 5, 6, 7},
		},
		{
			arr:        testArr,
			pattern:    "P|pictures",
			after:      0,
			before:     0,
			context:    0,
			ignoreCase: false,
			invert:     false,
			fixed:      true,
			result:     []int{7},
		},
		{
			arr:        testArr,
			pattern:    "virtualbox",
			after:      0,
			before:     0,
			context:    0,
			ignoreCase: true,
			invert:     false,
			fixed:      true,
			result:     []int{3, 4},
		},
	}

	for i, testCase := range data {
		args := arguments{
			after:      &testCase.after,
			before:     &testCase.before,
			context:    &testCase.context,
			count:      &testCase.count,
			ignoreCase: &testCase.ignoreCase,
			invert:     &testCase.invert,
			fixed:      &testCase.fixed,
			lineNum:    &testCase.lineNum,
		}
		res, err := findNumsOfLines(testArr, testCase.pattern, args)
		if !reflect.DeepEqual(res, testCase.result) {
			t.Errorf("test %v, Expected %v, got %v", i, testCase.result, res)
		}
		if err != nil {
			t.Error(err)
		}
	}
}

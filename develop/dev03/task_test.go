package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	testArr := [][]string{
		{"янв", "15", "Downloads", "5b"},
		{"дек", "6", "Android", "4"},
		{"июн", "10", "Sources", "45"},
		{"окт", "31", "VirtualBox", "90"},
		{"окт", "31", "VirtualBox", "90"},
		{"янв", "13", "Lightworks", "9k"},
		{"янв", "11", "Pictures", "12"},
	}
	data := []struct {
		arr    [][]string
		k      int
		n      bool
		r      bool
		u      bool
		result [][]string
		err    error
	}{
		{
			arr: testArr,
			k:   3,
			n:   false,
			r:   false,
			u:   false,
			result: [][]string{
				{"дек", "6", "Android", "4"},
				{"янв", "15", "Downloads", "5b"},
				{"янв", "13", "Lightworks", "9k"},
				{"янв", "11", "Pictures", "12"},
				{"июн", "10", "Sources", "45"},
				{"окт", "31", "VirtualBox", "90"},
				{"окт", "31", "VirtualBox", "90"},
			},
			err: nil,
		},
		{
			arr: testArr,
			k:   3,
			n:   false,
			r:   true,
			u:   false,
			result: [][]string{
				{"окт", "31", "VirtualBox", "90"},
				{"окт", "31", "VirtualBox", "90"},
				{"июн", "10", "Sources", "45"},
				{"янв", "11", "Pictures", "12"},
				{"янв", "13", "Lightworks", "9k"},
				{"янв", "15", "Downloads", "5b"},
				{"дек", "6", "Android", "4"},
			},
			err: nil,
		},
		{
			arr: testArr,
			k:   3,
			n:   false,
			r:   false,
			u:   true,
			result: [][]string{
				{"дек", "6", "Android", "4"},
				{"янв", "15", "Downloads", "5b"},
				{"янв", "13", "Lightworks", "9k"},
				{"янв", "11", "Pictures", "12"},
				{"июн", "10", "Sources", "45"},
				{"окт", "31", "VirtualBox", "90"},
			},
			err: nil,
		},
		{
			arr: testArr,
			k:   4,
			n:   false,
			r:   false,
			u:   false,
			result: [][]string{
				{"янв", "11", "Pictures", "12"},
				{"дек", "6", "Android", "4"},
				{"июн", "10", "Sources", "45"},
				{"янв", "15", "Downloads", "5b"},
				{"окт", "31", "VirtualBox", "90"},
				{"окт", "31", "VirtualBox", "90"},
				{"янв", "13", "Lightworks", "9k"},
			},
			err: nil,
		},
		{
			arr: testArr,
			k:   4,
			n:   true,
			r:   false,
			u:   false,
			result: [][]string{
				{"дек", "6", "Android", "4"},
				{"янв", "15", "Downloads", "5b"},
				{"янв", "13", "Lightworks", "9k"},
				{"янв", "11", "Pictures", "12"},
				{"июн", "10", "Sources", "45"},
				{"окт", "31", "VirtualBox", "90"},
				{"окт", "31", "VirtualBox", "90"},
			},
			err: nil,
		},
		{
			arr: testArr,
			k:   2,
			n:   true,
			r:   true,
			u:   true,
			result: [][]string{
				{"окт", "31", "VirtualBox", "90"},
				{"янв", "15", "Downloads", "5b"},
				{"янв", "13", "Lightworks", "9k"},
				{"янв", "11", "Pictures", "12"},
				{"июн", "10", "Sources", "45"},
				{"дек", "6", "Android", "4"},
			},
			err: nil,
		},
		{
			arr:    testArr,
			k:      0,
			n:      true,
			r:      true,
			u:      true,
			result: nil,
			err:    fmt.Errorf("invalid field number"),
		},
	}

	for i, testCase := range data {
		res, err := sort(testCase.k, testCase.n, testCase.r, testCase.u, testCase.arr)
		if !reflect.DeepEqual(res, testCase.result) {
			t.Errorf("test %v, Expected %v, got %v", i, testCase.result, res)
		}
		if err != nil && err.Error() != testCase.err.Error() {
			t.Errorf("Expected %+v, got %+v", testCase.err.Error(), err.Error())
		}
	}
}

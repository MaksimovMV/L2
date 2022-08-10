package main

import (
	"reflect"
	"testing"
)

func TestFindAnagram(t *testing.T) {
	arr := []string{"пятка", "пятак", "тяпка", "листок", "слиток", "пятка", "столик"}
	result := map[string][]string{
		"пятка":  {"пятак", "пятка", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
	}

	res := findAnagram(&arr)

	if !reflect.DeepEqual(*res, result) {
		t.Errorf("Expected %v, got %v", result, res)
	}
}

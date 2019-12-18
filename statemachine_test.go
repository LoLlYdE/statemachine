package main

import (
	"strings"
	"testing"
)

func TestPreParseLine(t *testing.T) {

	expected := strings.Split("This is the expected result", " ")
	got := preParseLine("This is the expected result // this is a test comment")
	if !stringSliceEquals(expected, got) {
		t.Errorf("parseData(\"\") failed, expected %v, got %v", expected, got)
	}

	got = preParseLine("This is the expected result // testing //multiple //comments // in // a // line")
	if !stringSliceEquals(expected, got) {
		t.Errorf("parseData(\"\") failed, expected %v, got %v", expected, got)
	}

	got = preParseLine("This is the expected result /// testing triple slash")
	if !stringSliceEquals(expected, got) {
		t.Errorf("parseData(\"\") failed, expected %v, got %v", expected, got)
	}

	got = preParseLine("This is the expected result")
	if !stringSliceEquals(expected, got) {
		t.Errorf("parseData(\"\") failed, expected %v, got %v", expected, got)
	}

	// dont forget to test empty strings
	expected = make([]string, 0)
	got = preParseLine("")
	if !stringSliceEquals(expected, got) {
		t.Errorf("parseData(\"\") failed, expected %v, got %v", expected, got)
	}
}

func TestParseData(t *testing.T) {
	// test for empty data
	emptyData := parseData("")

	if len(emptyData) > 0 {
		t.Errorf("parseData(\"\") failed, expected %v, got %v", make([]bool, 0), emptyData)
	}

	// test example data
	exampleData := parseData("11000101100101001100")
	exampleDataCorrect := []bool{true, true, false, false, false, true, false, true, true, false, false, true, false, true, false, false, true, true, false, false}
	if !boolSliceEquals(exampleData, exampleDataCorrect) {
		t.Errorf("parseData(\"\") failed, expected %v, got %v", exampleDataCorrect, exampleData)
	}

}

func boolSliceEquals(data, correct []bool) bool {
	if len(data) != len(correct) {
		return false
	}
	for i, item := range data {
		if item != correct[i] {
			return false
		}
	}
	return true
}

func stringSliceEquals(data, correct []string) bool {
	if len(data) != len(correct) {
		return false
	}
	for i, item := range data {
		if item != correct[i] {
			return false
		}
	}
	return true
}

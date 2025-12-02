package main

import (
	"fmt"
	"slices"
	"testing"
)

const input1 = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestDay2Part1(t *testing.T) {
	result := Part1(input1)
	expected := 1227775554

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay2Part2(t *testing.T) {
	result := Part2(input1)
	expected := 4174379265

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestIsAnyRepeatedSequences(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123123", true},
		{"11", true},
		{"111", true},
		{"1188511885", true},
		{"97", false},
		{"98", false},
		{"98", false},
		{"2121212124", false},
	}
	for _, test := range tests {
		result := IsAnyRepeatedSequences(test.input)
		if result != test.expected {
			fmt.Printf("For input %v, got %v, expected %v\n", test.input, result, test.expected)
			t.Fail()
		}
	}
}

func TestIsAnyRepeatedSequenceOptimizied(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"hello", false},
		{"hellohello", true},
	}
	for _, test := range tests {
		result := IsAnyRepeatedSequencesOptimized(test.input)
		if result != test.expected {
			fmt.Printf("For input %v, got %v, expected %v\n", test.input, result, test.expected)
			t.Fail()
		}
	}
}

func TestInvalidIdRanges(t *testing.T) {

	tests := []struct {
		input    string
		expected []int
	}{
		{"11-22", []int{11, 22}},
		{"95-115", []int{99, 111}},
		{"998-1012", []int{999, 1010}},
		{"1188511880-1188511890", []int{1188511885}},
		{"222220-222224", []int{222222}},
		{"1698522-1698528", []int{}},
		{"446443-446449", []int{446446}},
		{"38593856-38593862", []int{38593859}},
		{"565653-565659", []int{565656}},
		{"824824821-824824827", []int{824824824}},
		{"2121212118-2121212124", []int{2121212121}},
	}

	for _, test := range tests {
		result := GetInvalidIdsFromRange(test.input)
		for _, id := range test.expected {
			found := slices.Contains(result, id)
			if !found {
				fmt.Printf("For input %v, expected to find %v in result %v\n", test.input, id, result)
				t.Fail()
			}
		}
		for _, id := range result {
			found := slices.Contains(test.expected, id)
			if !found {
				fmt.Printf("For input %v, did not expect to find %v in result %v\n", test.input, id, result)
				t.Fail()
			}
		}
	}
}

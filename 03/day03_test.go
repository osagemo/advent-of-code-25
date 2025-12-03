package main

import (
	"fmt"
	"testing"
)

const input1 = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestDay3Part1(t *testing.T) {
	result := Part1(input1)
	expected := 357

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestMaxJoltageFromTwoBatteriesForBank(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}

	for _, test := range tests {
		result := MaxJoltageFromBatteriesInBank(test.input, 2)
		if result != test.expected {
			fmt.Printf("For input %v, got %v, expected %v\n", test.input, result, test.expected)
			t.Fail()
		}
	}
}

func TestMaxJoltageFromTwelveBatteriesForBank(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}

	for _, test := range tests {
		result := MaxJoltageFromBatteriesInBank(test.input, 12)
		if result != test.expected {
			fmt.Printf("For input %v, got %v, expected %v\n", test.input, result, test.expected)
			t.Fail()
		}
	}
}

func TestDay3Part2(t *testing.T) {
	result := Part2(input1)
	expected := 3121910778619

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

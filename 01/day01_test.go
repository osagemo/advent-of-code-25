package main

import (
	"fmt"
	"testing"
)

const input1 = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

const inputTestPositive = `R50
R60
R40
R500
R99
R1`

const inputTestNegative = `R50
L1
R1
L1
R1`

func TestDay1Part1(t *testing.T) {
	result := Part1(input1)
	expected := 3

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay1Part2Positive(t *testing.T) {
	result := Part2(inputTestPositive)
	expected := 8

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay1Part2Negative(t *testing.T) {
	result := Part2(inputTestNegative)
	expected := 3

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay1Part2(t *testing.T) {
	result := Part2(input1)
	expected := 6

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

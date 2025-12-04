package main

import (
	"fmt"
	"testing"

	"github.com/osagemo/advent-of-code/utils"
)

const input1 = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestDay4Part1(t *testing.T) {
	utils.Debug = true
	result := Part1(input1)
	expected := 13

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay4Part2(t *testing.T) {
	result := Part2(input1)
	expected := 43

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

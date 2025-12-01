package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/osagemo/advent-of-code-24/internal/utils"
)

//go:embed input.txt
var input string

// How many times does the dial end up at 0?
func Part1(input string) int {
	dialPosition := 50
	instructions := strings.Split(input, "\n")
	dialEndsAtZeroCount := 0

	for _, instruction := range instructions {
		direction, distance := parseInstruction(instruction)
		switch direction {
		case "R":
			dialPosition = (dialPosition + distance) % 100
		case "L":
			dialPosition = (dialPosition - distance) % 100
			// Negative remainder
			if dialPosition < 0 {
				dialPosition += 100
			}
		}
		if dialPosition == 0 {
			dialEndsAtZeroCount++
		}
	}

	return dialEndsAtZeroCount
}

// How many times does the dial pass 0?
func Part2(input string) int {
	dialPosition := 50
	dialZeroCount := 0
	instructions := strings.Split(input, "\n")

	for _, instruction := range instructions {
		// position + distance / 100
		direction, distance := parseInstruction(instruction)
		zeroCount := 0
		switch direction {
		case "R":
			zeroCount = (dialPosition + distance) / 100
			dialPosition = (dialPosition + distance) % 100
		case "L":
			if distance >= dialPosition && dialPosition > 0 {
				// We are going below zero, how many full rotations after?
				zeroCount = 1 + (distance-dialPosition)/100
			} else if dialPosition == 0 {
				// how many full rotations?
				zeroCount = distance / 100
			}

			dialPosition = (dialPosition - distance) % 100
			if dialPosition < 0 {
				dialPosition += 100
			}
		}
		dialZeroCount += zeroCount
	}

	return dialZeroCount
}

func parseInstruction(instruction string) (direction string, distance int) {
	direction = (string)(instruction[0])
	distance = utils.MustParseInt(instruction[1:])
	if direction != "R" && direction != "L" {
		panic("Malformed input. Accepted directions are 'R' and 'L'")
	}
	return
}

func main() {
	input := strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.Trim(input, "\n")
	fmt.Println("Day 1")
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input))
	fmt.Println(time.Since(start))
}

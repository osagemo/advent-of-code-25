package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/osagemo/advent-of-code/utils"
)

//go:embed input.txt
var input string

const paper = '@'

// How many rolls of paper can be accessed by a forklift?
func Part1(input string) int {
	matrix := parseMatrix(input)
	accessiblePaperRolls := findAccessiblePaperRolls(matrix)

	return len(accessiblePaperRolls)
}

// Once a roll of paper is removed, the forklifts might be able to access more rolls of paper,
// which they might also be able to remove.
// How many total rolls of paper could the Elves remove if they keep repeating this process?
func Part2(input string) int {
	matrix := parseMatrix(input)
	totalPapersRemoved := 0

	// quick hack reusing part 1 implementation, todo: optimize
	papersToRemove := findAccessiblePaperRolls(matrix)
	iterations := 0
	for len(papersToRemove) > 0 {
		iterations++
		totalPapersRemoved += len(papersToRemove)
		for _, coord := range papersToRemove {
			matrix[coord.Row][coord.Col] = '.'
		}
		papersToRemove = findAccessiblePaperRolls(matrix)
	}
	utils.DebugLog("\nn iterations: %v\n", iterations)
	return totalPapersRemoved
}

// The forklifts can only access a roll of paper if there are fewer than four rolls of paper in the eight adjacent positions
func findAccessiblePaperRolls(matrix [][]rune) []utils.Coordinate {
	start := utils.Coordinate{Row: 0, Col: 0}
	coordinatesToCheck := []utils.Coordinate{start}
	checked := map[utils.Coordinate]bool{start: true}
	accessiblePaperRolls := []utils.Coordinate{}

	for len(coordinatesToCheck) > 0 {
		// pop
		coord := coordinatesToCheck[0]
		coordinatesToCheck = coordinatesToCheck[1:]

		paperNeighbours := 0
		for _, dir := range utils.AllDirections {
			neighbour := coord.Add(dir)
			if utils.WithinBounds(matrix, neighbour) {
				if !checked[neighbour] {
					coordinatesToCheck = append(coordinatesToCheck, neighbour)
					checked[neighbour] = true
				}
				if matrix[neighbour.Row][neighbour.Col] == paper {
					paperNeighbours++
				}
			}
		}

		if matrix[coord.Row][coord.Col] == paper && paperNeighbours < 4 {
			accessiblePaperRolls = append(accessiblePaperRolls, coord)
		}
	}

	utils.DebugLog("found rolls: %v\n", accessiblePaperRolls)

	printMatrix(matrix, accessiblePaperRolls)
	return accessiblePaperRolls
}

func parseMatrix(input string) [][]rune {
	lines := strings.Split(input, "\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	return matrix
}

func printMatrix(matrix [][]rune, accessiblePapers []utils.Coordinate) {
	for rowIndex, row := range matrix {
		for ColIndex, c := range row {
			if slices.Contains(accessiblePapers, utils.Coordinate{Row: rowIndex, Col: ColIndex}) {
				utils.DebugLog("x")
			} else {
				utils.DebugLog("%s", (string)(c))
			}
		}
		utils.DebugLog("\n")
	}
}

func main() {
	input := utils.PrepareInput(input)
	fmt.Println("Day 4")
	utils.RunTimed("Part 1", func() int { return Part1(input) })
	utils.RunTimed("Part 2", func() int { return Part2(input) })
}

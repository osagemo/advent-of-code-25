package main

import (
	_ "embed"
	"fmt"
	"strconv"

	"strings"

	"github.com/osagemo/advent-of-code/utils"
)

//go:embed input.txt
var input string

func Part1(input string) int {
	totalJoltage := 0
	batteryBanks := strings.SplitSeq(input, "\n")
	for bank := range batteryBanks {
		totalJoltage += MaxJoltageFromBatteriesInBank(bank, 2)
	}
	return totalJoltage
}

func Part2(input string) int {
	totalJoltage := 0
	batteryBanks := strings.SplitSeq(input, "\n")
	for bank := range batteryBanks {
		totalJoltage += MaxJoltageFromBatteriesInBank(bank, 12)
	}
	return totalJoltage
}

func MaxJoltageFromBatteriesInBank(bank string, nBatteries int) int {
	utils.DebugLog("\n Bank: %s\n", bank)

	selected := make([]int, nBatteries)
	batteries := make([]int, len(bank))
	for i, b := range bank {
		batteries[i], _ = strconv.Atoi(string(b))
	}

	maxUsedIndex := -1
	for pos := range selected {
		maxIndexForPos := -1
		searchBoundary := len(batteries) - nBatteries + pos

		for i := maxUsedIndex + 1; i <= searchBoundary; i++ {
			if maxIndexForPos == -1 || batteries[i] > batteries[maxIndexForPos] {
				maxIndexForPos = i
			}
		}
		selected[pos] = batteries[maxIndexForPos]
		maxUsedIndex = maxIndexForPos
	}

	utils.DebugLog("%v\n", batteries)
	return utils.ConcatDigitsToInt(selected...)
}

func main() {
	input := utils.PrepareInput(input)
	fmt.Println("Day 3")
	utils.RunTimed("Part 1", func() int { return Part1(input) })
	utils.RunTimed("Part 2", func() int { return Part2(input) })
}

package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/osagemo/advent-of-code/utils"
)

//go:embed input.txt
var input string

// Which Ids in the ranges provided as input consist of a sequence of digits repeated twice?
func Part1(input string) int {
	productIdRanges := strings.SplitSeq(input, ",")
	invalidIdSum := 0

	for r := range productIdRanges {
		fields := strings.Split(r, "-")
		if len(fields) != 2 {
			panic("invalid range")
		}
		from := utils.MustParseInt(fields[0])
		to := utils.MustParseInt(fields[1])

		for id := from; id <= to; id++ {

			if IsRepeatedSequence(strconv.Itoa(id)) {
				invalidIdSum += id
				utils.DebugLog("%v is invalid, adding to sum\n", id)
			}
		}

	}
	return invalidIdSum
}

// Which Ids in the ranges provided as input consist of a sequence of digits repeated n times?
func Part2(input string) int {
	productIdRanges := strings.SplitSeq(input, ",")
	invalidIdSum := 0

	for r := range productIdRanges {
		invalidIds := GetInvalidIdsFromRange(r)
		for _, id := range invalidIds {
			invalidIdSum += id
		}
	}
	return invalidIdSum
}

func GetInvalidIdsFromRange(idRange string) []int {
	invalidIds := []int{}
	utils.DebugLog("\nProcessing range: %v\n", idRange)
	fields := strings.Split(idRange, "-")
	if len(fields) != 2 {
		panic("invalid range")
	}
	from := utils.MustParseInt(fields[0])
	to := utils.MustParseInt(fields[1])

	for id := from; id <= to; id++ {
		if IsAnyRepeatedSequences(strconv.Itoa(id)) {
			invalidIds = append(invalidIds, id)
			utils.DebugLog("%v is invalid, adding to sum\n", id)
		}
	}

	return invalidIds
}

// Couldn't figure it out myself. From discussions about similar leetcode problem https://leetcode.com/problems/repeated-substring-pattern/description/

// "A string consists of the same pattern repeated multiple times if and only if the string is a nontrivial rotation of itself."
func IsAnyRepeatedSequencesOptimized(id string) bool {
	// "If x and y are strings of the same length, then x is a rotation of y if and only if x is a substring of yy."

	// id: "1212"
	// (id + id): "12121212"
	// (id + id)[1:]: "2121212"
	// strings.Index(..., "1212"): 1
	// 1 + 1 != 4 -> true
	return strings.Index((id + id)[1:], id)+1 != len(id)
}

// Brutish, TODO: optimize
func IsAnyRepeatedSequences(id string) bool {
	utils.DebugLog("\nProcessing %s\n", id)
	length := len(id)

	for size := 1; size <= length/2; size++ {
		if length%size != 0 {
			continue
		}
		repeated := true

		// whole string consists of 0:size
		for i := 0; i <= length-size; i += size {
			if id[i:i+size] != id[0:size] {
				utils.DebugLog("%v is not like 0-%v: %v, break\n", id[i:i+size], size, id[0:size])
				repeated = false
				break
			}
		}
		if repeated {
			return true
		}
	}
	return false
}

func IsRepeatedSequence(id string) bool {
	if len(id)%2 != 0 {
		return false
	}

	half := len(id) / 2
	first, second := id[:half], id[half:]
	return first == second
}

// TODO: works in regex101 for all examples but go does not support capture groups in regexp patterns :( workaround?
// https://github.com/golang/go/issues/42549
func IsRepeatedSequenceRegex(id string) bool {
	// regex inspo: https://stackoverflow.com/a/2823037
	// The whole string needs to consist of the capture group exactly twice
	repeatedDigits, err := regexp.MatchString(`^([0-9]+)\1$`, id)
	if err != nil {
		panic(fmt.Sprintf("unable to compile regex pattern: %v", err))
	}

	return repeatedDigits
}

func main() {
	input := utils.PrepareInput(input)
	fmt.Println("Day 2")
	utils.RunTimed("Part 1", func() int { return Part1(input) })
	utils.RunTimed("Part 2", func() int { return Part2(input) })
}

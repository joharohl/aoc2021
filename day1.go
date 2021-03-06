package main

import (
	"fmt"
	"strconv"

	"github.com/joharohl/aoc2021/utils"
)

func strToInt(strs []string) (ints []int) {
	ints = make([]int, 0, len(strs))
	for _, str := range strs {
		integer, err := strconv.Atoi(str)
		utils.CheckError(err)
		ints = append(ints, integer)
	}
	return
}

func sumThree(ints []int) (sums []int) {
	sums = make([]int, 0, len(ints)/3)

	for i := 0; i < len(ints); i++ {
		if i+2 == len(ints) {
			break
		}
		sum := ints[i] + ints[i+1] + ints[i+2]
		sums = append(sums, sum)
	}
	return
}

func countIncreasingValues(ints []int) (count int) {
	count = 0
	for i := 1; i < len(ints); i++ {
		if ints[i] > ints[i-1] {
			count++
		}
	}
	return
}

func main() {
	measurements := strToInt(utils.ReadFile("inputs/day1.txt"))

	// Find increasing values
	fmt.Println(countIncreasingValues(measurements))

	// Find sliding window of increasing values
	fmt.Println(countIncreasingValues(sumThree(measurements)))

}

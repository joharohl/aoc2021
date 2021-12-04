package main

import (
	"fmt"
	"math/bits"
	"strconv"

	"github.com/joharohl/aoc2021/utils"
)

func strToInt(strs []string) (ints []uint) {
	ints = make([]uint, 0, len(strs))
	for _, str := range strs {
		integer, err := strconv.ParseInt(str, 2, 0)
		utils.CheckError(err)
		ints = append(ints, uint(integer))
	}
	return
}

func getBitX(integer uint, x int) uint {
	return (integer & (1 << x)) >> x
}

func getBitXList(integers []uint, x int) (result []uint) {
	result = make([]uint, len(integers))
	for _, integer := range integers {
		result = append(result, getBitX(integer, x))
	}
	return
}

func getMostCommonBitValue(integers []uint, x int) string {
	numberOfMeasurements := len(integers)
	sumColum0 := sum(getBitXList(integers, x))
	if sumColum0*2 >= numberOfMeasurements {
		return "1"
	} else {
		return "0"
	}
}

func getMostCommonBitValues(integers []uint) (result string) {
	nColumns := bits.Len(integers[0])
	resultSlice := make([]string, nColumns)
	for i := 0; i < nColumns; i++ {
		resultSlice[nColumns-1-i] = getMostCommonBitValue(integers, i)
	}
	for _, v := range resultSlice {
		result = result + v
	}
	return
}

func findOxygenGeneratorRating(integers []uint) uint {
	nColumns := bits.Len(integers[0])
	candidates := make([]uint, len(integers))
	copy(candidates, integers)
	for i := nColumns - 1; i >= 0; i-- {
		candidates = filterMostCommon(candidates, i)
		if len(candidates) == 1 {
			break
		}
	}
	return candidates[0]
}

func findCO2ScrubberRating(integers []uint) uint {
	nColumns := bits.Len(integers[0])
	candidates := make([]uint, len(integers))
	copy(candidates, integers)
	for i := nColumns - 1; i >= 0; i-- {
		candidates = filterLeastCommon(candidates, i)
		if len(candidates) == 1 {
			break
		}
	}
	return candidates[0]
}

func filterMostCommon(integers []uint, column int) (out []uint) {
	out = make([]uint, 0, len(integers))
	mostCommonBitValue, err := strconv.ParseInt(getMostCommonBitValue(integers, column), 2, 0)
	utils.CheckError(err)
	for _, v := range integers {
		if ((v >> column) & 1) == uint(mostCommonBitValue) {
			out = append(out, v)
		}
	}
	return
}

func filterLeastCommon(integers []uint, column int) (out []uint) {
	out = make([]uint, 0, len(integers))
	mostCommonBitValue, err := strconv.ParseInt(getMostCommonBitValue(integers, column), 2, 0)
	leastCommonBitValue := mostCommonBitValue ^ 1
	utils.CheckError(err)
	for _, v := range integers {
		if ((v >> column) & 1) == uint(leastCommonBitValue) {
			out = append(out, v)
		}
	}
	return
}

func sum(array []uint) int {
	result := 0
	for _, v := range array {
		result += int(v)
	}
	return result
}

func printBitSlices(integers []uint) {
	fmt.Println("-----")
	for _, v := range integers {
		fmt.Printf("%05b\n", v)
	}
	fmt.Println("-----")
}

func main() {
	measurements := strToInt(utils.ReadFile("inputs/day3.txt"))
	gammaRateTmp, err := strconv.ParseInt(getMostCommonBitValues(measurements), 2, 0)
	utils.CheckError(err)
	gammaRate := uint(gammaRateTmp)
	epsilonRate := gammaRate ^ 0b111111111111

	fmt.Printf("%012b %02d\n", gammaRate, gammaRate)
	fmt.Printf("%012b %02d\n", epsilonRate, epsilonRate)

	fmt.Println(epsilonRate * gammaRate)

	oxygenGeneratorRating := findOxygenGeneratorRating(measurements)
	co2ScrubberRating := findCO2ScrubberRating(measurements)
	fmt.Println(oxygenGeneratorRating * co2ScrubberRating)

}

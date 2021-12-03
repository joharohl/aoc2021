package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joharohl/aoc2021/utils"
)

func main() {
	movements := utils.ReadFile("inputs/day2.txt")

	depth := 0
	distance := 0
	aim := 0

	for _, movement := range movements {
		command := strings.Split(movement, " ")
		value, err := strconv.Atoi(command[1])
		utils.CheckError(err)

		switch command[0] {
		case "forward":
			distance += value
			depth += value * aim
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}
	fmt.Println(aim * distance)
	fmt.Println(depth * distance)

}

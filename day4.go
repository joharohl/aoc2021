package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joharohl/aoc2021/utils"
)

var Reset = "\033[0m"
var Bold = "\033[1m"

type board struct {
	numbers    []string
	checked    []string
	size       int
	hasBingoed bool
}

func newBoard(numbers []string) *board {
	b := board{numbers: numbers, size: 5}
	return &b
}

func (b *board) CheckNumber(numbers ...string) {
	for _, n := range numbers {
		b.checked = append(b.checked, n)
	}
}

func (b *board) isAllChecked(numbers []string) bool {
	all := true
	for _, number := range numbers {
		all = all && b.isChecked(number)
	}
	return all
}

func (b *board) isChecked(number string) (isChecked bool) {
	isChecked = false
	for _, n := range b.checked {
		if n == number {
			isChecked = true
		}
	}
	return
}

func (b *board) getColum(x int) (numbers []string) {
	numbers = make([]string, b.size, b.size)
	for i := 0; i < b.size; i++ {
		numbers[i] = b.numbers[x+i*b.size]
	}
	return
}

func (b *board) getRow(y int) (numbers []string) {
	numbers = make([]string, b.size, b.size)
	for i := 0; i < b.size; i++ {
		numbers[i] = b.numbers[y*b.size+i]
	}
	return
}

func (b *board) PrintScore(number string) {
	sum := 0
	n, err := strconv.Atoi(number)
	utils.CheckError(err)
	for _, number := range b.numbers {
		if !b.isChecked(number) {
			i, err := strconv.Atoi(number)
			utils.CheckError(err)
			sum += i
		}
	}
	fmt.Printf("Score: %d\n", sum*n)
}

func (b *board) bingo() (isBingo bool) {
	isBingo = false
	for i := 0; i < b.size; i++ {
		any := false
		any = any || b.isAllChecked(b.getColum(i))
		any = any || b.isAllChecked(b.getRow(i))
		if any {
			isBingo = true
			break
		}
	}
	return
}

func (b *board) Print() {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			number := b.numbers[i*b.size+j]
			if b.isChecked(number) {
				number = Bold + number + Reset
			}
			fmt.Printf("% 2s ", number)
		}
		fmt.Printf("\n")
	}
}

func main() {
	var drawings []string
	boards := make([]*board, 0, 50)
	file, err := os.Open("inputs/day4.txt")
	utils.CheckError(err)
	scanner := bufio.NewScanner(file)
	i := -1
	var tmpNumbers []string

	for scanner.Scan() {
		i++
		if i == 0 {
			drawings = strings.Split(scanner.Text(), ",")
			continue
		}

		if scanner.Text() == "" {
			if len(tmpNumbers) == 25 {
				fmt.Println("Saving board X")
				boards = append(boards, newBoard(tmpNumbers))
			}
			fmt.Println("Start of new board")
			tmpNumbers = make([]string, 0, 25)
			continue
		}
		tmpNumbers = append(tmpNumbers, strings.Fields(scanner.Text())...)
	}
	if len(tmpNumbers) == 25 {
		fmt.Println("Saving board X")
		boards = append(boards, newBoard(tmpNumbers))
	}
	utils.CheckError(scanner.Err())

	nBoardsBingoed := 0
	for _, drawing := range drawings {
		for _, board := range boards {
			board.CheckNumber(drawing)
			if !board.hasBingoed && board.bingo() {
				nBoardsBingoed++
				if nBoardsBingoed == 1 {
					board.Print()
					board.PrintScore(drawing)
				}
				if nBoardsBingoed == len(boards) {
					board.Print()
					board.PrintScore(drawing)
				}
				board.hasBingoed = true
			}
		}
		if nBoardsBingoed == len(boards) {
			break
		}

	}

}

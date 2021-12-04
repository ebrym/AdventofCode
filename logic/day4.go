package logic

import (
	"adventofcode/general"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Day4Task1() int {
	log.Println("Day 4")

	input, err := general.ReadLines("assets/day4.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("Part 4:", input)

	var numbers []int
	for _, s := range strings.Split(input[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)

	var boards [][]int

	for i := 2; i < len(input); i += 6 {
		var board []int
		for _, s := range strings.Split(strings.Join(input[i:i+5], " "), " ") {
			if s == "" {
				continue
			}

			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("readLines: %s", err)
			}
			board = append(board, i)
		}
		if len(board) != 25 {
			log.Fatal(board)
		}
		boards = append(boards, board)
	}

	fmt.Println(boards)

	for _, n := range numbers {
		for _, b := range boards {
			for i, v := range b {
				if v == n {
					b[i] = 0
					break
				}
			}

			// checking win
			if checkWin(b) {
				fmt.Println(n, b)
				sum := 0
				for _, j := range b {
					sum += j
				}
				//fmt.Println(n, b)
				//fmt.Println(sum * n)

				return sum * n
			}
		}
	}
	return 0
}
func Day4Task2() int {
	log.Println("Day 4")

	input, err := general.ReadLines("assets/day4.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("Part 4:", input)

	var numbers []int
	for _, s := range strings.Split(input[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)

	var boards [][]int

	for i := 2; i < len(input); i += 6 {
		var board []int
		for _, s := range strings.Split(strings.Join(input[i:i+5], " "), " ") {
			if s == "" {
				continue
			}

			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("readLines: %s", err)
			}
			board = append(board, i)
		}
		if len(board) != 25 {
			log.Fatal(board)
		}
		boards = append(boards, board)
	}

	// keeping track of the boards
	boardWin := make([]bool, len(boards))

	for _, n := range numbers {
		for b := range boards {
			if boardWin[b] {
				continue
			}

			//boardWin[b] := false
			for i, v := range boards[b] {
				if v == n {
					boards[b][i] = 0
					break
				}
			}

			// checking win
			if checkWin(boards[b]) {
				sum := 0
				for _, j := range boards[b] {
					sum += j
				}
				fmt.Println(n, boards[b], sum*n)
				boardWin[b] = true
				//break
			}
		}
	}
	return 0
}

func checkWin(b []int) bool {

	win := true
	for i := 0; i < 5; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 5; i < 10; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 10; i < 15; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 15; i < 20; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 20; i < 25; i++ {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}
	//=========================
	win = true
	for i := 0; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 1; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 2; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 3; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	win = true
	for i := 4; i < 25; i += 5 {
		if b[i] != 0 {
			win = false
		}
	}
	if win {
		return true
	}

	return false
}

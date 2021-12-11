package logic

import (
	"fmt"
	"log"
)

// type Point struct {
// 	x1, x2 int
// 	y1, y2 int
// }

func Day11Task1() int {
	log.Println("Day 10 task 1")

	// data, err := general.ReadLines("assets/day10.txt")
	// if err != nil {
	// 	log.Fatalf("readLines: %s", err)
	// }

	//fmt.Println("Day 5 Part 1:", input)

	// data := []string{
	// 	"5483143223",
	// 	"2745854711",
	// 	"5264556173",
	// 	"6141336146",
	// 	"6357385478",
	// 	"4167524645",
	// 	"2176841721",
	// 	"6882881134",
	// 	"4846848554",
	// 	"5283751526",
	// }

	data := []string{
		"5212166716",
		"1567322581",
		"2268461548",
		"3481561744",
		"6248342248",
		"6526667368",
		"5627335775",
		"8124511754",
		"4614137683",
		"4724561156",
	}

	var grid [][]int = make([][]int, 10)

	for i := range grid {
		grid[i] = make([]int, 10)

		for j := 0; j < 10; j++ {
			grid[i][j] = int(data[i][j] - byte('0'))
		}
	}
	count := 0
	// for i := 0; i < 100; i++ {
	// 	count += checkStep(grid)
	// }

	syncCount := 0
	for i := 0; i < 1000; i++ {
		thisCount := checkStep(grid)
		syncCount += thisCount
		if thisCount == 100 {
			fmt.Println("SYNC", i+1)
			break
		}
	}
	return count
}

func checkStep(grid [][]int) int {
	flashed := [10][10]bool{}

	//1
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x]++
		}
	}

	//1

	updatedSomething := true

	for updatedSomething {
		updatedSomething = false
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if flashed[y][x] {
					continue
				}
				if grid[y][x] > 9 {

					updatedSomething = true
					flashed[y][x] = true

					if y-1 >= 0 {
						if x-1 >= 0 {
							grid[y-1][x-1]++
						}
						grid[y-1][x]++
						if x+1 < len(grid[y]) {
							grid[y-1][x+1]++
						}
					}

					if x-1 >= 0 {
						grid[y][x-1]++
					}
					if x+1 < len(grid[y]) {
						grid[y][x+1]++
					}

					if y+1 < len(grid) {
						if x-1 >= 0 {
							grid[y+1][x-1]++
						}
						grid[y+1][x]++
						if x+1 < len(grid[y]) {
							grid[y+1][x+1]++
						}
					}
				}
			}
		}
		// repeat
	}

	count := 0
	for y := 0; y < len(flashed); y++ {
		for x := 0; x < len(flashed[y]); x++ {
			if flashed[y][x] {
				grid[y][x] = 0
				count++
			}
		}
	}
	return count
}

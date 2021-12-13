package logic

import (
	"adventofcode/general"
	"fmt"
	"log"
)

type xy struct {
	x, y int
}

func Day13Task1() int {
	log.Println("Day 13 task")

	input, err := general.ReadLines("assets/day13.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// visited
	// path tracking
	// nodes
	data := input
	// data := []string{
	// 	"start-A",
	// 	"start-b",
	// 	"A-c",
	// 	"A-b",
	// 	"b-d",
	// 	"A-end",
	// 	"b-end",
	// }
	var points []xy
	var folds []xy

	parsefold := false

	for _, row := range data {
		if row == "" {
			parsefold = true
			continue
		}

		if !parsefold {
			var xy xy
			fmt.Sscanf(row, "%d,%d", &xy.x, &xy.y)
			points = append(points, xy)
		} else {
			var dir rune
			var value int
			fmt.Sscanf(row, "fold along %c=%d", &dir, &value)
			if dir == 'x' {
				folds = append(folds, xy{x: value})
			} else {
				folds = append(folds, xy{y: value})
			}
		}
	}

	fmt.Println(points)
	fmt.Println(folds)
	//fmt.Println(caves["b"])

	maxX := folds[0].x * 2
	maxY := folds[1].y * 2

	var grid [][]int = make([][]int, maxY+1)
	for i := range grid {
		grid[i] = make([]int, maxX+1)
	}

	count := 0
	for _, p := range points {

		var x, y int = p.x, p.y

		if p.x > folds[0].x {
			x = 2*folds[0].x - p.x
		}

		if grid[y][x] == 0 {
			count++
			grid[y][x] = 1
		}
	}
	fmt.Println(count)

	return 0
}
func Day13Task2() int {
	log.Println("Day 13 task 2")

	input, err := general.ReadLines("assets/day13.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// visited
	// path tracking
	// nodes
	data := input
	var points []xy
	var folds []xy

	parsefold := false

	for _, row := range data {
		if row == "" {
			parsefold = true
			continue
		}

		if !parsefold {
			var xy xy
			fmt.Sscanf(row, "%d,%d", &xy.x, &xy.y)
			points = append(points, xy)
		} else {
			var dir rune
			var value int
			fmt.Sscanf(row, "fold along %c=%d", &dir, &value)
			if dir == 'x' {
				folds = append(folds, xy{x: value})
			} else {
				folds = append(folds, xy{y: value})
			}
		}
	}

	//fmt.Println(points)
	//fmt.Println(folds)
	//fmt.Println(caves["b"])

	//maxX := folds[0].x * 2
	//maxY := folds[1].y * 2

	var grid [][]int = make([][]int, 6)
	for i := range grid {
		grid[i] = make([]int, 40)
	}

	for _, p := range points {

		var x, y int = p.x, p.y
		for _, f := range folds {
			// for x fold
			if f.x > 0 {
				if x > f.x {
					x = 2*f.x - x
				}
			} else {
				if y > f.y {
					y = 2*f.y - y
				}
			}
		}
		grid[y][x] = 1
	}
	for _, row := range grid {

		//fmt.Println(row)
		for _, v := range row {
			if v == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	return 0
}

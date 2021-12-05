package logic

import (
	"adventofcode/general"
	"fmt"
	"log"
	"math"
)

type Point struct {
	x1, x2 int
	y1, y2 int
}

func Day5Task1() int {
	log.Println("Day 5")

	input, err := general.ReadLines("assets/day5.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("Day 5 Part 1:", input)

	var entries = make([]Point, len(input))
	vhLines := []Point{}
	maxX := 0
	maxY := 0

	for i, l := range input {
		fmt.Sscanf(l, "%d,%d -> %d,%d",
			&entries[i].x1, &entries[i].y1,
			&entries[i].x2, &entries[i].y2)

		if entries[i].x1 > maxX {
			maxX = entries[i].x1
		}
		if entries[i].x2 > maxX {
			maxX = entries[i].x2
		}
		if entries[i].y1 > maxY {
			maxX = entries[i].y1
		}
		if entries[i].y2 > maxY {
			maxX = entries[i].y2
		}

		if entries[i].x1 == entries[i].x2 || entries[i].y1 == entries[i].y2 {
			vhLines = append(vhLines, entries[i])
		}

	}
	fmt.Println(entries)

	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}

	for _, l := range vhLines {
		if l.x1 == l.x2 {
			y1 := l.y1
			y2 := l.y2
			if y2 < y1 {
				y1 = l.y2
				y2 = l.y1
			}
			for y := y1; y <= y2; y++ {
				grid[y][l.x1]++
			}
		} else {
			x1 := l.x1
			x2 := l.x2
			if x2 < x1 {
				x1 = l.x2
				x2 = l.x1
			}
			for x := x1; x <= x2; x++ {
				grid[l.y1][x]++
			}
		}

	}

	count := 0
	for _, row := range grid {
		for _, v := range row {
			if v >= 2 {
				count++
			}
		}
	}
	//fmt.Println("Part 4:", input)

	return count
}
func Day5Task2() int {
	log.Println("Day 5")

	input, err := general.ReadLines("assets/day5.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("Day 5 Part 1:", input)

	var entries = make([]Point, len(input))
	vhLines := []Point{}
	maxX := 0
	maxY := 0

	for i, l := range input {
		fmt.Sscanf(l, "%d,%d -> %d,%d",
			&entries[i].x1, &entries[i].y1,
			&entries[i].x2, &entries[i].y2)

		if entries[i].x1 > maxX {
			maxX = entries[i].x1
		}
		if entries[i].x2 > maxX {
			maxX = entries[i].x2
		}
		if entries[i].y1 > maxY {
			maxX = entries[i].y1
		}
		if entries[i].y2 > maxY {
			maxX = entries[i].y2
		}

		if entries[i].x1 == entries[i].x2 || entries[i].y1 == entries[i].y2 {
			vhLines = append(vhLines, entries[i])
		}

	}
	fmt.Println(entries)

	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}

	for _, l := range entries {
		if l.x1 == l.x2 {
			y1 := l.y1
			y2 := l.y2
			if y2 < y1 {
				y1 = l.y2
				y2 = l.y1
			}
			for y := y1; y <= y2; y++ {
				grid[y][l.x1]++
			}
		} else if l.y1 == l.y2 {
			x1 := l.x1
			x2 := l.x2
			if x2 < x1 {
				x1 = l.x2
				x2 = l.x1
			}
			for x := x1; x <= x2; x++ {
				grid[l.y1][x]++
			}
		} else { // for 45 degree lines
			dx := int(float64(l.x2-l.x1) / math.Abs(float64(l.x2-l.x1)))
			dy := int(float64(l.y2-l.y1) / math.Abs(float64(l.y2-l.y1)))
			x := l.x1
			y := l.y1
			for y != l.y2+dy {
				grid[y][x]++
				x += dx
				y += dy
			}
		}

	}

	count := 0
	for _, row := range grid {
		for _, v := range row {
			if v >= 2 {
				count++
			}
		}
	}
	//fmt.Println("Part 4:", input)

	return count
}

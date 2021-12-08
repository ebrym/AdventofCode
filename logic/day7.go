package logic

import (
	"adventofcode/general"
	"log"
	"math"
	"strconv"
	"strings"
)

// type Point struct {
// 	x1, x2 int
// 	y1, y2 int
// }

func Day7Task1() int {
	log.Println("Day 7 task 1")

	input, err := general.ReadLines("assets/day7.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	//fmt.Println("Day 5 Part 1:", input)

	var pos []int
	max := 0
	for _, v := range strings.Split(input[0], ",") {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		if i > max {
			max = i
		}
		pos = append(pos, i)
	}

	var dist []int = make([]int, max+1)
	for i := 0; i <= max; i++ {
		for _, v := range pos {
			dist[i] += int(math.Abs(float64(v - i)))
		}
	}

	min := 99999999
	for _, v := range dist {
		if v < min {
			min = v
		}
	}
	return min
}

func Day7Task2() int {
	log.Println("Day 7 task 2")

	input, err := general.ReadLines("assets/day7.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	//fmt.Println("Day 5 Part 1:", input)

	var pos []int
	max := 0
	for _, v := range strings.Split(input[0], ",") {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		if i > max {
			max = i
		}
		pos = append(pos, i)
	}

	var dist []int = make([]int, max+1)
	for i := 0; i <= max; i++ {
		for _, v := range pos {
			dist[i] += cost(int(math.Abs(float64(v - i))))
		}
	}

	min := 999999999
	for _, v := range dist {
		if v < min {
			min = v
		}
	}
	return min
}
func cost(i int) int {
	return (i * (i + 1)) / 2
}

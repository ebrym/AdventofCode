package logic

import (
	"adventofcode/general"
	"log"
	"strconv"
	"strings"
)

// type Point struct {
// 	x1, x2 int
// 	y1, y2 int
// }

func Day6Task1() int {
	log.Println("Day 6 task 1")

	input, err := general.ReadLines("assets/day6.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	//fmt.Println("Day 5 Part 1:", input)

	var fish []int
	for _, n := range strings.Split(input[0], ",") {
		i, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		fish = append(fish, i)
	}
	//var sum int = 0
	for i := 0; i < 80; i++ {
		fish = step(fish)
	}

	// for _, v := range fish {
	// 	sum += v
	// }
	//fmt.Println(step(fish))
	return len(fish)
}
func step(fish []int) []int {
	var newfish []int
	for i := range fish {
		if fish[i] == 0 {
			newfish = append(newfish, 8)
			fish[i] = 6
		} else {
			fish[i]--
		}
	}
	return append(fish, newfish...)
}
func Day6Task2() int {
	log.Println("Day 6 task 2")

	input, err := general.ReadLines("assets/day6.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	//fmt.Println("Day 5 Part 1:", input)

	var fish = make([]int, 9)
	for _, n := range strings.Split(input[0], ",") {
		i, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		fish[i]++
	}

	var sum int = 0
	for i := 0; i < 256; i++ {
		fish = step2(fish)
	}
	for _, v := range fish {
		sum += v
	}
	//fmt.Println(step(fish))
	return sum
}
func step2(fish []int) []int {
	var next = make([]int, 9)
	for i := 1; i < 9; i++ {
		next[i-1] = fish[i]
	}
	next[6] += fish[0]
	next[8] += fish[0]
	return next
}

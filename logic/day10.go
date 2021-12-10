package logic

import (
	"adventofcode/general"
	"fmt"
	"log"
	"sort"
)

// type Point struct {
// 	x1, x2 int
// 	y1, y2 int
// }

func Day10Task1() int {
	log.Println("Day 10 task 1")

	input, err := general.ReadLines("assets/day10.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	//fmt.Println("Day 5 Part 1:", input)

	completion := []string{}
	//count := 0
	score := 0
	for _, line := range input {
		stack := []rune{}
		corrupted := rune(0)

	loop:
		for _, c := range line {
			switch c {
			case '(', '[', '{', '<':
				stack = append(stack, c)
			case ')', ']', '}', '>':
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if c == ')' && pop == '(' ||
					c == ']' && pop == '[' ||
					c == '}' && pop == '{' ||
					c == '>' && pop == '<' {
					continue
				} else {
					corrupted = c
					break loop
				}
			}
		}

		switch corrupted {
		case ')':
			score += 3
		case ']':
			score += 57
		case '}':
			score += 1197
		case '>':
			score += 25137
		default:
			completion = append(completion, string(stack))
		}
	}

	log.Println("Incomplete:", completion)

	//total := 0
	scores := make([]int, len(completion))
	for i, c := range completion {
		for j := len(c) - 1; j >= 0; j-- {
			scores[i] *= 5

			switch c[j] {
			case '(':
				scores[i] += 1
			case '[':
				scores[i] += 2
			case '{':
				scores[i] += 3
			case '<':
				scores[i] += 4
			}
		}

	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
	return score
}

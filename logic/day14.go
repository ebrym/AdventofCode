package logic

import (
	"adventofcode/general"
	"fmt"
	"log"
)

func Day14Task1() int {
	log.Println("Day 14 task")

	input, err := general.ReadLines("assets/day14.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// visited
	// path tracking
	// nodes
	// data := []string{
	// 	"start-A",
	// 	"start-b",
	// 	"A-c",
	// 	"A-b",
	// 	"b-d",
	// 	"A-end",
	// 	"b-end",
	// }

	data := input

	chain := data[0]

	// fmt.Println(chain)
	rules := map[string]string{}
	for _, rule := range data[2:] {
		var k, v string
		fmt.Sscanf(rule, "%s -> %s", &k, &v)
		rules[k] = v
	}

	fmt.Println(rules)
	fmt.Println(chain)

	var next string = chain
	for i := 0; i < 40; i++ {
		next = readSteps(next, rules)
	}

	fmt.Println(next)
	fmt.Println(countSteps(next))

	//task 1 result calculation
	//4156-1038=3118
	return 0
}

func Day14Task2() int {
	log.Println("Day 14 task 2")

	input, err := general.ReadLines("assets/day14.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	data := input

	chain := data[0]

	rules := map[string]Output{}
	for _, rule := range data[2:] {
		var k, v string
		fmt.Sscanf(rule, "%s -> %s", &k, &v)

		rules[k] = Output{
			Pair1:      string(k[0]) + v,
			Pair2:      string(k[1]) + v,
			NewElement: v,
		}
	}

	s := newStep(chain)
	// var next string = chain
	for i := 0; i < 40; i++ {
		s = newReadStep(s, rules)
	}

	fmt.Println(s.count)
	min := 99999999999999
	max := 0
	for _, v := range s.count {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	// fmt.Println(countSteps(next))

	fmt.Println(min, max, max-min)
	//4156-1038=3118
	return 0
}

type Step struct {
	pairs map[string]int
	count map[string]int
}
type Output struct {
	Pair1, Pair2 string
	NewElement   string
}

func newStep(in string) *Step {
	s := &Step{
		pairs: map[string]int{},
		count: map[string]int{},
	}
	for i := 0; i < len(in)-1; i++ {
		pair := in[i : i+2]
		s.pairs[pair]++
	}
	for i := 0; i < len(in); i++ {
		s.count[string(in[i])]++
	}

	return s

}
func newReadStep(s *Step, rules map[string]Output) *Step {
	out := &Step{
		pairs: map[string]int{},
		count: map[string]int{},
	}

	for k, v := range s.count {
		out.count[k] = v
	}

	for pair, count := range s.pairs {
		t := rules[pair]
		out.pairs[t.Pair1] += count
		out.pairs[t.Pair2] += count
		out.count[t.NewElement] += count
	}
	return out
}
func readSteps(in string, rules map[string]string) string {
	next := ""
	for i := 0; i < len(in)-1; i++ {
		pair := in[i : i+2]
		ins := rules[pair]
		next += pair[:1] + ins
	}
	next += in[len(in)-1:]
	return next
}
func countSteps(in string) map[string]int {
	out := map[string]int{}
	for _, c := range in {
		out[string(c)]++
	}
	return out
}

package logic

import (
	"container/heap"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var inputFile14 = flag.String("inputFile15", "assets/day15.txt", "Relative file path to use as input.")

//var tr = otel.Tracer("day15")

type Coords struct {
	R, C int
}
type RiskMap map[Coords]int
type HeapQueue struct {
	Elems     *[]Coords
	Score     RiskMap
	Positions map[Coords]int
}

func (h HeapQueue) Len() int           { return len(*h.Elems) }
func (h HeapQueue) Less(i, j int) bool { return h.Score[(*h.Elems)[i]] < h.Score[(*h.Elems)[j]] }
func (h HeapQueue) Swap(i, j int) {
	h.Positions[(*h.Elems)[i]], h.Positions[(*h.Elems)[j]] = h.Positions[(*h.Elems)[j]], h.Positions[(*h.Elems)[i]]
	(*h.Elems)[i], (*h.Elems)[j] = (*h.Elems)[j], (*h.Elems)[i]
}

func (h HeapQueue) Push(x interface{}) {
	h.Positions[x.(Coords)] = len(*h.Elems)
	*h.Elems = append(*h.Elems, x.(Coords))
}

func (h HeapQueue) Pop() interface{} {
	old := *h.Elems
	n := len(old)
	x := old[n-1]
	*h.Elems = old[0 : n-1]
	delete(h.Positions, x)
	return x
}

func (h HeapQueue) Position(x Coords) int {
	if pos, ok := h.Positions[x]; ok {
		return pos
	}
	return -1
}

func Day15Task1() int {
	log.Println("Day 15 task")

	// input, err := general.ReadLines("assets/day15.txt")
	// if err != nil {
	// 	log.Fatalf("readLines: %s", err)
	// }

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

	flag.Parse()

	bytes, err := ioutil.ReadFile(*inputFile14)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]

	risks := make(RiskMap)
	for r, line := range split {
		for c, risk := range line {
			risks[Coords{r, c}] = int(risk - '0')
		}
	}
	height := len(split)
	width := len(split[0])
	start := Coords{0, 0}
	dst := Coords{height - 1, width - 1}
	fmt.Println(AStar(risks, &start, &dst))

	dst = Coords{height*5 - 1, width*5 - 1}
	expandedRisks := make(RiskMap)
	for k, v := range risks {
		for r := 0; r < 5; r++ {
			for c := 0; c < 5; c++ {
				increase := r + c
				value := 1 + (v+increase-1)%9
				expandedRisks[Coords{k.R + r*height, k.C + c*width}] = value
			}
		}
	}
	fmt.Println(AStar(expandedRisks, &start, &dst))

	//task 1 result calculation
	//4156-1038=3118
	return 0
}

func AStar(r RiskMap, src, dst *Coords) int {
	gScore := RiskMap{
		*src: 0,
	}
	fScore := RiskMap{
		*src: src.Heuristic(dst),
	}
	workList := HeapQueue{&[]Coords{*src}, fScore, make(RiskMap)}
	heap.Init(&workList)
	history := make(map[Coords]Coords)

	for len(*workList.Elems) != 0 {
		// Pop the current node off the worklist.
		current := heap.Pop(&workList).(Coords)

		if current == *dst {
			// Reconstruct the score by retracing our path to start.
			score := 0
			for current != *src {
				score += r[current]
				current = history[current]
			}
			return score
		}
		for _, n := range r.Neighbors(current) {
			proposedScore := gScore[current] + r[n]
			if previousScore, ok := gScore[n]; !ok || proposedScore < previousScore {
				history[n] = current
				gScore[n] = proposedScore
				fScore[n] = proposedScore + n.Heuristic(dst)
				if pos := workList.Position(n); pos == -1 {
					heap.Push(&workList, n)
				} else {
					heap.Fix(&workList, pos)
				}
			}
		}
	}
	return -1
}

func (c Coords) Heuristic(dst *Coords) int {
	// Manhattan distance, assuming min of 1 per traverse.
	return (dst.C - c.C) + (dst.R - c.R)
}

func (r RiskMap) Neighbors(pos Coords) []Coords {
	var coords []Coords
	up := Coords{pos.R - 1, pos.C}
	down := Coords{pos.R + 1, pos.C}
	left := Coords{pos.R, pos.C - 1}
	right := Coords{pos.R, pos.C + 1}
	for _, v := range []Coords{up, down, left, right} {
		if _, ok := r[v]; ok {
			coords = append(coords, v)
		}
	}
	return coords
}

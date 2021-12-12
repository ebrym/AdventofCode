package logic

import (
	"adventofcode/general"
	"log"
	"strings"
)

type Cave struct {
	Name  string
	Big   bool
	Other []string
}

func Day12Task1() int {
	log.Println("Day 12 task")

	input, err := general.ReadLines("assets/day12.txt")
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

	caves := map[string]*Cave{}

	for _, edge := range data {
		var sp = strings.Split(edge, "-")
		var a, b = sp[0], sp[1]

		if _, ok := caves[a]; !ok {
			caves[a] = &Cave{
				Name:  a,
				Big:   a[0] < 'a',
				Other: []string{b},
			}
		} else {
			caves[a].Other = append(caves[a].Other, b)
		}

		a, b = b, a
		if _, ok := caves[a]; !ok {
			caves[a] = &Cave{
				Name:  a,
				Big:   a[0] < 'a',
				Other: []string{b},
			}
		} else {
			caves[a].Other = append(caves[a].Other, b)
		}
	}
	path := navigate(caves)
	//fmt.Println(caves["b"])
	return len(path)
}

type Path struct {
	Path              []string
	Visited           map[string]bool
	VisitedSmallTwice bool
}

func copyMap(a map[string]bool) map[string]bool {
	out := map[string]bool{}
	for k, v := range a {
		out[k] = v
	}
	return out
}

func navigate(caves map[string]*Cave) []string {

	var queue = []Path{Path{[]string{"start"}, map[string]bool{}, false}}
	var paths = []string{}

	for len(queue) > 0 {
		var cur = queue[0]
		queue = queue[1:]

		cave := caves[cur.Path[len(cur.Path)-1]]

		if cave.Name == "end" {
			paths = append(paths, strings.Join(cur.Path, ","))
			continue
		}

		newVisited := copyMap(cur.Visited)
		if !cave.Big {
			newVisited[cave.Name] = true
		}

		for _, cave := range cave.Other {
			newVisitedSmallTwice := cur.VisitedSmallTwice
			if cur.Visited[cave] {
				if cave == "start" || cur.VisitedSmallTwice {
					continue
				} else {
					newVisitedSmallTwice = true
				}
			}
			newPath := make([]string, len(cur.Path))
			copy(newPath, cur.Path)
			newPath = append(newPath, cave)
			queue = append(queue, Path{newPath, newVisited, newVisitedSmallTwice})
		}
	}
	return paths
}

package logic

import (
	"adventofcode/general"
	"fmt"
	"log"
	"sort"
	"strings"
)

// type Point struct {
// 	x1, x2 int
// 	y1, y2 int
// }

func Day8Task1() int {
	log.Println("Day 8 task 1")

	input, err := general.ReadLines("assets/day8.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	//fmt.Println("Day 5 Part 1:", input)

	count := 0
	for _, v := range input {
		v2 := strings.Split(v, " | ")
		v3 := strings.Split(v2[1], " ")

		for _, d := range v3 {
			if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
				count++
			}
		}
	}

	return count
}

func Day8Task2() int {
	log.Println("Day 8 task 2")

	input, err := general.ReadLines("assets/day8.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	//fmt.Println("Day 5 Part 1:", input)

	count := 0
	sum := 0
	for _, v := range input {
		v2 := strings.Split(v, " | ")
		v3 := strings.Split(v2[1], " ")

		for _, d := range v3 {
			if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
				count++
			}
		}

		digits := check(strings.Split(v2[0], " "))

		output := digits[normal(v3[0])]*1000 +
			digits[normal(v3[1])]*100 +
			digits[normal(v3[2])]*10 +
			digits[normal(v3[3])]

		sum += output

	}
	fmt.Println(count, sum)
	return count
}
func check(num []string) map[string]int {
	var len2, len3, len4, len7 string
	var len5 []string
	var len6 []string
	for _, d := range num {
		if len(d) == 2 {
			len2 = normal(d)
		} else if len(d) == 3 {
			len3 = normal(d)
		} else if len(d) == 4 {
			len4 = normal(d)
		} else if len(d) == 7 {
			len7 = normal(d)
		} else if len(d) == 5 {
			len5 = append(len5, normal(d))
		} else {
			len6 = append(len6, normal(d))
		}
	}

	out := map[string]int{
		len2: 1,
		len4: 4,
		len3: 7,
		len7: 8,
	}
	var digit6 string
	// for 6
	for i, d := range len6 {
		if strings.Index(d, string(len2[0])) == -1 {
			out[d] = 6
			digit6 = d
			len6 = append(len6[:i], len6[i+1:]...)
			break
		}
		if strings.Index(d, string(len2[1])) == -1 {
			out[d] = 6
			digit6 = d
			len6 = append(len6[:i], len6[i+1:]...)
			break
		}
	}

	// for not 4 should be 0
	for i, d := range len6 {
		if strings.Index(d, string(len4[0])) == -1 {
			out[d] = 0
			len6 = append(len6[:i], len6[i+1:]...)
			break
		}
		if strings.Index(d, string(len4[1])) == -1 {
			out[d] = 0
			len6 = append(len6[:i], len6[i+1:]...)
			break
		}
		if strings.Index(d, string(len4[2])) == -1 {
			out[d] = 0
			len6 = append(len6[:i], len6[i+1:]...)
			break
		}
		if strings.Index(d, string(len4[3])) == -1 {
			out[d] = 0
			len6 = append(len6[:i], len6[i+1:]...)
			break
		}
	}

	// for 9
	out[len6[0]] = 9

	// for 3
	for i, d := range len5 {
		if strings.Index(d, string(len2[0])) != -1 &&
			strings.Index(d, string(len2[1])) != -1 {
			out[d] = 3
			len5 = append(len5[:i], len5[i+1:]...)
			break
		}
	}

	// find 5 for all in 6
	for i, d := range len5 {
		if strings.Index(digit6, string(d[0])) != -1 &&
			strings.Index(digit6, string(d[1])) != -1 &&
			strings.Index(digit6, string(d[2])) != -1 &&
			strings.Index(digit6, string(d[3])) != -1 &&
			strings.Index(digit6, string(d[4])) != -1 {
			out[d] = 5
			len5 = append(len5[:i], len5[i+1:]...)
			break
		}
	}

	// for 2
	out[len5[0]] = 2
	return out
}

func normal(digit string) string {
	arr := strings.Split(digit, "")
	sort.Strings(arr)
	return strings.Join(arr, "")
}

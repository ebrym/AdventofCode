package logic

// --- Day 3: Power! ---
// The submarine has been making some odd creaking noises, so you ask it to produce a diagnostic report just in case.

// The diagnostic report (your puzzle input) consists of a list of binary numbers which, when decoded properly, can tell you many useful things about the conditions of the submarine. The first parameter to check is the power consumption.

// You need to use the binary numbers in the diagnostic report to generate two new binary numbers (called the gamma rate and the epsilon rate). The power consumption can then be found by multiplying the gamma rate by the epsilon rate.

// Each bit in the gamma rate can be determined by finding the most common bit in the corresponding position of all numbers in the diagnostic report. For example, given the following diagnostic report:

// 00100
// 11110
// 10110
// 10111
// 10101
// 01111
// 00111
// 11100
// 10000
// 11001
// 00010
// 01010
// Considering only the first bit of each number, there are five 0 bits and seven 1 bits. Since the most common bit is 1, the first bit of the gamma rate is 1.

// The most common second bit of the numbers in the diagnostic report is 0, so the second bit of the gamma rate is 0.

// The most common value of the third, fourth, and fifth bits are 1, 1, and 0, respectively, and so the final three bits of the gamma rate are 110.

// So, the gamma rate is the binary number 10110, or 22 in decimal.

// The epsilon rate is calculated in a similar way; rather than use the most common bit, the least common bit from each position is used. So, the epsilon rate is 01001, or 9 in decimal. Multiplying the gamma rate (22) by the epsilon rate (9) produces the power consumption, 198.

// Use the binary numbers in your diagnostic report to calculate the gamma rate and epsilon rate, then multiply them together. What is the power consumption of the submarine? (Be sure to represent your answer in decimal, not binary.)

import (
	"adventofcode/general"
	"log"
	"strconv"
	"strings"
)

func Day3Task1() int64 {

	log.Println("Day 2")

	input, err := general.ReadLines("assets/day3.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// result
	var result int64 = 0

	var nums []int = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < len(input); i++ {
		var chars = strings.Split(input[i], "")
		for j := 0; j < len(chars); j++ {
			var a int = 0
			a, _ = strconv.Atoi(chars[j])
			nums[j] += a
		}
	}
	var gamma, epsilon string = "", ""
	for i := 0; i < len(nums); i++ {
		if nums[i] > 500 {
			gamma += "1"
			epsilon += "0"
		}
		if nums[i] < 500 {
			gamma += "0"
			epsilon += "1"
		}
	}
	var g, e int64 = 0, 0
	g, _ = strconv.ParseInt(gamma, 2, 64)
	e, _ = strconv.ParseInt(epsilon, 2, 64)
	result = g * e
	return result
}

// func findCommon(list []string, index int, invert bool = false, ox bool = true) string {

// 	var result int = 0
// 	var items []int = []int{0,0}
// 	for i := 0; i < len(list); i++ {
// 		var a string = strconv.Itoa(input[i])
// 		chars[i] = a
// 	}

// 	var chars []string = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}
// 	for i := 0; i < len(input); i++ {
// 		var a string = strconv.Itoa(input[i])
// 		chars[i] = a
// 	}
// 	var s string = strings.Join(chars, "")
// 	var chars2 []string = strings.Split(s, "")
// 	var map1 map[string]int = make(map[string]int)
// 	var map2 map[string]int = make(map[string]int)
// 	for i := 0; i < len(chars2); i++ {
// 		map1[chars2[i]]++
// 	}
// 	for i := 0; i < len(chars2); i++ {
// 		map2[chars2[i]]++
// 	}
// 	for k, v := range map1 {
// 		if map2[k] > v {
// 			result = v
// 		}
// 	}
// 	return result
// }

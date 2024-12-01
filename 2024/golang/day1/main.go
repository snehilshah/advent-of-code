package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"golang/utils"
)

//go:embed input.txt
var inputDay string

func parse(input string) (left []int, right []int) {
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)

		var l, _ = strconv.Atoi(fields[0])
		var r, _ = strconv.Atoi(fields[1])

		left = append(left, l)
		right = append(right, r)
	}

	return
}

func Part1(input string) int {
	var left, right = parse(input)
	slices.Sort(left)
	slices.Sort(right)

	var res int
	for i := 0; i < len(left); i++ {
		res += utils.Abs(left[i] - right[i])
	}
	return res
}

func Part2(input string) int {
	var left, right = parse(input)
	var occurrence = make(map[int]int)
	for _, v := range right {
		occurrence[v]++
	}
	var res int
	for _, v := range left {
		n, ok := occurrence[v]
		if ok && n > 0 {
			res += n * v
		}
	}
	return res
}

func main() {
	fmt.Println("--2024 day 01 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))
	fmt.Println("part2", Part2(inputDay))
}

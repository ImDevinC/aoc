package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const input = "input.txt"

func main() {
	payload, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	list := string(payload)
	list1, list2 := parseList(list)
	part1(list1, list2)
	part2(list1, list2)
}

func part2(list1 []int, list2 []int) {
	var score int
	for _, x := range list1 {
		count := 0
		for _, y := range list2 {
			if x == y {
				count++
			}
		}
		score += x * count
	}
	fmt.Println("Score: ", score)
}

func part1(list1 []int, list2 []int) {
	var dist int
	for i := range len(list1) {
		if list1[i] > list2[i] {
			dist += list1[i] - list2[i]
		} else {
			dist += list2[i] - list1[i]
		}
	}
	fmt.Println("Total Distance: ", dist)
}

func parseList(list string) ([]int, []int) {
	first := []int{}
	second := []int{}

	for _, s := range strings.Split(list, "\n") {
		if len(s) == 0 {
			continue
		}
		result := strings.Split(s, " ")
		f, _ := strconv.Atoi(string(result[0]))
		s, _ := strconv.Atoi(string(result[3]))
		first = append(first, f)
		second = append(second, s)
	}
	sort.Ints(first)
	sort.Ints(second)

	return first, second
}

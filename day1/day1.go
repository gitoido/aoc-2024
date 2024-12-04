package main

import (
	"aoc/shared"
	"fmt"
	"log"
	"sort"
)

type PairOfLists struct {
	ListA []int
	ListB []int
}

func (l PairOfLists) getDistance() int {
	if len(l.ListA) != len(l.ListB) {
		log.Fatalf("len(listA) != len(listB)\n")
		return 0
	}

	sort.Slice(l.ListA, func(i, j int) bool {
		return l.ListA[i] < l.ListA[j]
	})

	sort.Slice(l.ListB, func(i, j int) bool {
		return l.ListB[i] < l.ListB[j]
	})

	distance := 0
	for i := 0; i < len(l.ListA); i++ {
		distance = distance + shared.IntDifference(l.ListA[i], l.ListB[i])
	}
	return distance
}

func (l PairOfLists) SimilarityScore() int {
	dictA := make(map[int]int)
	dictB := make(map[int]int)

	for _, v := range l.ListA {
		dictA[v]++
	}
	for _, v := range l.ListB {
		dictB[v]++
	}

	var similarity int
	for key, _ := range dictA {
		if _, ok := dictB[key]; ok {
			similarity += key * dictB[key]
		}
	}

	return similarity
}

func main() {
	data := shared.ReadFileToIntSlice("day1/day1.txt")
	var p PairOfLists
	for i := 0; i < len(data); i++ {
		p.ListA = append(p.ListA, data[i][0])
		p.ListB = append(p.ListB, data[i][1])
	}

	fmt.Println(p.getDistance())
	fmt.Println(p.SimilarityScore())
}

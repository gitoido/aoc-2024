package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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
		if l.ListA[i] != l.ListB[i] {
			mx := math.Max(float64(l.ListA[i]), float64(l.ListB[i]))
			mn := math.Min(float64(l.ListA[i]), float64(l.ListB[i]))
			distance = distance + int(mx-mn)
		}
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
	f, err := os.OpenFile("day1/day1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var p PairOfLists
	for sc.Scan() {
		items := strings.Fields(sc.Text())
		a, _ := strconv.Atoi(items[0])
		b, _ := strconv.Atoi(items[1])
		p.ListA = append(p.ListA, a)
		p.ListB = append(p.ListB, b)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}

	fmt.Println(p.getDistance())
	fmt.Println(p.SimilarityScore())
}

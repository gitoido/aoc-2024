package shared

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReadFileToIntSlice(path string) [][]int {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
		return nil
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var items [][]int
	for sc.Scan() {
		stringSlice := strings.Fields(sc.Text())
		ints := make([]int, len(stringSlice))

		for i, s := range stringSlice {
			ints[i], _ = strconv.Atoi(s)
		}

		items = append(items, ints)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return nil
	}

	return items
}

func IntDifference(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func CheckSorting(s []int, r bool) bool {
	res := true

	if r {
		slices.Reverse(s)
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] >= s[i+1] {
			res = false
			break
		}
	}

	return res
}

func RemoveSliceElement(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

package shared

import (
	"bufio"
	"log"
	"os"
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

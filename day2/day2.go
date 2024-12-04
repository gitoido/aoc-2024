package main

import (
	"aoc/shared"
	"fmt"
)

type ReportList [][]int

func (l ReportList) ValidCount() int {
	count := 0
	for i := 0; i < len(l); i++ {
		sort := shared.CheckSorting(l[i], false)
		sortReverse := shared.CheckSorting(l[i], true)
		safe := checkSafety(l[i])

		if safe != true {
			for j := 0; j < len(l[i]); j++ {
				cropped := shared.RemoveSliceElement(l[i], j)
				safe = checkSafety(cropped)
				if safe {
					break
				}
			}
		}

		if safe && (sort || sortReverse) {
			count++
		}
	}
	return count
}

func checkSafety(l []int) bool {
	safe := true
	for j := 0; j < len(l); j++ {
		if j+1 > len(l)-1 {
			continue
		}

		distance := shared.IntDifference(l[j], l[j+1])
		if distance < 1 || distance > 3 {
			safe = false
			break
		}
	}
	return safe
}

func main() {
	var reports ReportList
	data := shared.ReadFileToIntSlice("day2/day2.txt")

	for i := 0; i < len(data); i++ {
		reports = append(reports, data[i])
	}

	fmt.Println(reports.ValidCount())
}

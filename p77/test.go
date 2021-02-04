package main

import (
	"fmt"
	"sort"
)

func overlappingList(l [][]int) (nl [][]int) {
	nl = append(nl, l[0])
	for i := 1; i < len(l); i++ {
		nlLast := len(nl) - 1
		if nl[nlLast][1] < l[i][0] {
			nl = append(nl, l[i])
		} else if nl[nlLast][1] < l[i][1] {
			nl[nlLast][1] = l[i][1]
		}
	}
	return nl
}

func main() {
	l := [][]int{
		{1, 3},
		{5, 8},
		{4, 10},
		{20, 25},
	}
	fmt.Printf("input: %v \n", l)
	sort.SliceStable(l, func(i, j int) bool { return l[i][0] < l[j][0] })
	fmt.Printf("sort list: %v \n", l)

	nl := overlappingList(l)
	fmt.Printf("overlapping list: %v \n", nl)
}

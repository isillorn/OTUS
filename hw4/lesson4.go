package main

import (
	"fmt"
	"reflect"
)

// Concat slices
func Concat(slices [][]int) []int {
	var s []int
	for _, slice := range slices {
		s = append(s, slice...)
	}
	return s
}

func main() {
	type pair struct {
		s [][]int
		r []int
	}

	test := []pair{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{{1, 2}, {3, 4}, {6, 5}}, []int{1, 2, 3, 4, 6, 5}},
		{[][]int{{1, 2}, {}, {6, 5}}, []int{1, 2, 6, 5}},
	}
	for _, t := range test {
		s := t.s
		r := t.r
		r2 := Concat(s)
		fmt.Printf("Test: %v %v... ", s, r2)
		if ok := reflect.DeepEqual(r, r2); ok {
			fmt.Printf("%v\n", ok)
		}
	}
}

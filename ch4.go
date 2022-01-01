package main

import (
	"fmt"
	"math"
)


func main() {
	fmt.Printf("Deduping")
	slice := []int{1,4,3,44,43,5,3,2,1,32,24,5}
	fmt.Printf("%v -> ", slice)
	fmt.Printf("%v\n", dedup(slice))
	fmt.Printf("Flattening")
	nest := [][]int{{1, 2, 3, 4},
					 {5, 6, 7, 8},
					 {9, 10, 11},
					 {12, 13, 14, 15},
					 {16, 17, 18, 19, 20}}
	fmt.Printf("%v -> ", nest)
	fmt.Printf("%v\n", flatten(nest))
	fmt.Printf("2D-ing")
	twod := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Printf("%v -> \n", twod)
	fmt.Printf("%v\n", make2d(twod, 2))
}

func dedup(dup []int) (deduped []int) {
	valcount := make(map[int]int)
	for _, val := range dup {
		if valcount[val] < 1 {
			deduped = append(deduped, val)
			valcount[val] += 1
		}
	}
	return deduped
}

func flatten(nested [][]int) (flat []int) {
	for _, val := range nested {
		for _, inner_val := range val {
			flat = append(flat, inner_val)
		}
	}
	return flat
}

func make2d(flat []int, columns int) ([][]int) {
	rows := int(math.Ceil(float64(len(flat))/float64(columns)))
	nested := make([][]int, rows)
	for k := 0; k < rows; k++ {
		nested[k] = make([]int, columns)
	}
	for index, val := range flat {
		i := math.Floor(float64(index)/float64(columns))
		j := index % columns
		nested[int(i)][int(j)] = val
	}
	return nested
}




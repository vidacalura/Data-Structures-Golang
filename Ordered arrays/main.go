package main

import (
	"fmt"
	"sort"
)

// O(n)
func linear_search(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}

	return false
}

// O(log n)
func binary_search(arr []int, value int) bool {
	midpoint := len(arr) / 2

	if midpoint == 0 && value != arr[midpoint] {
		return false
	}

	if value < arr[midpoint] {
		return binary_search(arr[:midpoint], value)
	} else if value > arr[midpoint] {
		return binary_search(arr[midpoint + 1:], value)
	}

	return true
}

func main() {
	arr := [8]int{ 7, 4, 12, 6, 11, 9, 14, 16 }
	sort.Ints(arr[:])

	//fmt.Println(linear_search(arr[:], 14)) // true
	//fmt.Println(linear_search(arr[:], 17)) // false
	//fmt.Println(binary_search(arr[:], 7)) // true
	//fmt.Println(binary_search(arr[:], 8)) // false
}
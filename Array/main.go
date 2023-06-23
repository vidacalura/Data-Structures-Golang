// Módulo de arrays
// Cap. 1 - A Common-Sense Guide to Data Structures and Algorithms
package main

import (
	"fmt"
)

// Reading: O(n)
func read_array(arr []string) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

// Search: O(n)
func index_of(arr []string, value string) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}

	return -1
}

// Insertion: O(n+1)
func insert_at(arr []string, index int, value string) []string {
	if (index > len(arr)) { // impossível
		return arr
	} else if (index == len(arr)) { // best case scenario
		return append(arr, value)
	}

	arr = append(arr, "")
	copy(arr[index + 1:], arr[index:])
	arr[index] = value

	return arr
}

// Deletion: O(n+1)
func delete_at(arr []string, index int) []string {
	if (index >= len(arr)) {
		return arr
	}

	arr = append(arr, "")
	copy(arr[index:], arr[index + 1:])

	return arr[:len(arr) - 1]
}

func main() {
	var array []string = []string{ "apples", "bananas", "cucumbers", "dates", "elderberries" }

	//read_array(array)
	//fmt.Println(index_of(array, "apples"))
	//fmt.Println(index_of(array, "figs"))
	//fmt.Println(insert_at(array, 3, "figs"))
	//fmt.Println(delete_at(array, 3))
}
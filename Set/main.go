// Sets
// Cap. 1 - A Common-Sense Guide to Data Structures and Algorithms
package main

import (
	"fmt"
)

// Insertion: O(1)
func insert(set map[string]bool, value string) map[string]bool {
	set[value] = true
	return set
}

// Reading: O(n)
func read_set(set map[string]bool) {
	for key, _ := range set {
		fmt.Println(key)
	}
}

// Search: O(n)
func index_of(set map[string]bool, value string) bool {
	for key, _ := range set {
		if key == value {
			return true
		}
	}

	return false
}

func main() {
	// Golang não sets tem por padrão, por isso, precisamos
	// criar na mão esta estrutura de dados
	set := make(map[string]bool)

	set = insert(set, "teste")
	//read_set(set)
	//fmt.Println(index_of(set, "teste"))
	//delete(set, "teste") // Deletion: O(n+1)
}
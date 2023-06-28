// Implementação de Bubble Sort
package main

// O(n²)
func bubble_sort(arr []int) {
	for i, _ := range arr {
		for j, v := range arr[:len(arr)-1-i] {
			if v > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			} 
		}
	}
}
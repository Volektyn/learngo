// 1. Define a slice with unsorted numbers in it.
// 2. Print this slice to the console.
// 3. Sort the values using swapping.
// 4. Once done, print the now sorted numbers to the console.

package main

import "fmt"

func main() {
	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Println("Before: ", nums)
	BubbleSort(&nums)
	fmt.Println("After: ", nums)
}

func BubbleSort(nums *[]int) {
	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < len(*nums); i++ {
			if (*nums)[i-1] > (*nums)[i] {
				(*nums)[i], (*nums)[i-1] = (*nums)[i-1], (*nums)[i]
				swapped = true
			}
		}
	}
}

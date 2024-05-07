package main

import "fmt"

// It adds one element in each iteration.
// You need to select the smallest element in the array and move it to the beginning of the array by swapping with the front element
func main() {
	a := []int{-1, 10, 2, 9, 1, 9, 3, 6}
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		temp := a[i]
		a[i] = a[min]
		a[min] = temp
	}
	fmt.Println(a)
}

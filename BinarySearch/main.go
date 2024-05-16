package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := BinarySearch(a, 120)
	fmt.Println(index)

}

func BinarySearch(a []int, target int) int {
	low := 0
	high := len(a) - 1
	fmt.Println("low: ", low, "high: ", high)
	for low <= high {
		mid := low + ((high - low) / 2)
		fmt.Println("mid: ", mid)
		if target > a[mid] {
			low = mid + 1
		} else if target < a[mid] {
			high = mid - 1
		} else if target == a[mid] {
			return mid
		}

	}
	return -1
}

//TimeComplexity
//min - O(1) - if target ==a[mid] only i iteration
//max- O(logn) - maxium iteration are height of the tree -

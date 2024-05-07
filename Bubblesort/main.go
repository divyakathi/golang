package main

import "fmt"

func main() {
	//Two loops- 1. outer loop (i) is for comparisions
	//           2. Inner loop (j) is for swapping
	//In each iteration highest numbers will swap. If perform 1 iteration 1 highest number will sort,
	//                                             If we do 2 iterations , 2 highest numbers will sort..so on...

	//BubbleSort
	nums := []int{14, 5, 6, 8, 9}
	for i := 0; i < len(nums)-1; i++ {
		flag := false
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println(nums)
	// Time Complexity - O(n*n) - 2 for loops
}

/*
********1st highest number in an array*********
(Update the outer loop (i) with how many highest numbers you want )
nums := []int{16, 10, 6, 8, 5, -1, 0, 2}
	for i := 0; i < 1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}

	}
	fmt.Println(nums)

*******2 highest numbers in an array*********
	nums := []int{16, 10, 6, 8, 5, -1, 0, 2}
		for i := 0; i < 2; i++ {
			for j := 0; j < len(nums)-1-i; j++ {
				if nums[j] > nums[j+1] {
					temp := nums[j]
					nums[j] = nums[j+1]
					nums[j+1] = temp
				}
			}

		}
		fmt.Println(nums)


*/

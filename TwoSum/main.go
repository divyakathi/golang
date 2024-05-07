package main

import "fmt"

func main() {

	numbers := []int{2, 7, 11, 15, 11, 2, -6}
	target := 9
	output := make([][]int, 0)

	//Procedure 1:
	checkMap := make(map[int]int, 0)
	for i, v := range numbers {
		complement := target - v
		if val, ok := checkMap[complement]; ok {
			output = append(output, []int{i, val})
			continue
		}
		checkMap[v] = i
	}
	fmt.Println(checkMap)
	fmt.Println(output)

	/******************************************************************************
	//Procedure 2:
	checkmap := make(map[int]int, 0)
	for i, v := range numbers {
		checkmap[v] = i
	}

	for i := 0; i < len(numbers); i++ {
		complement := target - numbers[i]
		if _, ok := checkmap[complement]; ok {
			output = append(output, []int{i, checkmap[complement]})
		}
	}
	fmt.Println(output)
	*/
	//***********************************************************************************
	//Procedure 3:
	/*
		for i := 0; i < len(numbers); i++ {
			for j := i + 1; j < len(numbers); j++ {
				if target == numbers[i]+numbers[j] {
					output = append(output, []int{i, j})
				}
			}
		}
		fmt.Println(output) */
}

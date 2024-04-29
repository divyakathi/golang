package main

import "fmt"

func main() {
	numbers := []int{5, 5, 7, 2, 9, 4, 1}
	output := []int{}
	checkMap := make(map[int]int, 0)
	for k, v := range numbers {
		if checkMap[v] > 0 {
			checkMap[v] += 1
			continue
		}
		checkMap[v] = 1
		output = append(output, numbers[k])
	}
	fmt.Println("No of occurances : ", checkMap)
	fmt.Println("Non repitive numbers : ", output)

}

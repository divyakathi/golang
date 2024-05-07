package main

import "fmt"

func main() {
	//InsertionSort
	a := []int{10, 1, 15, 10, 12, 11, 13, 5, 6}
	for i := 1; i < len(a); i++ {
		j := i - 1
		temp := a[i]
		fmt.Println("Value of i,j,temp: ", i, j, temp)
		for j >= 0 && a[j] > temp {
			fmt.Println(i, "The value of j: ", j)
			fmt.Println(i, "   the numbers of j and temp : ", a[j], temp, a[j] > temp)
			a[j+1] = a[j]
			fmt.Println("Array--", a)
			j--

		}
		fmt.Println(i, "---OUTER LOOP----", j)
		a[j+1] = temp
		fmt.Println("updated array--", a)
	}

	fmt.Println(a)

}

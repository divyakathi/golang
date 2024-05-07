package main

import "fmt"

func main() {
	//InsertionSort
	a := []int{5, 2, 1, 9, 3, 6}
	/* for i := 1; i < len(a); i++ {
		fmt.Println("Value of i : ", i)
		j := i - 1
		temp := a[i]
		for j >= 0 {
			fmt.Println("value of i, j : ", i, j, temp)
			if a[j] > temp {
				fmt.Println("The value of j: ", j)
				fmt.Println(i, "   the numbers of j and i : ", a[j], a[i], temp, a[j] > temp)
				a[j+1] = a[j]
				fmt.Println("Array--", a)
			}
			j--
		}
		fmt.Println(i, "---OUTER LOOP----", j)
		a[j+1] = temp

		fmt.Println("updated array--", a)
	}
	*/
	fmt.Println(a)

}

package main

import "fmt"

func carAssemblyLineScheduling(a, ct [2][4]int, e, x [2]int) int {
	// T1 and T2 arrays to store minimum time taken on each line
	T1 := [4]int{}
	T2 := [4]int{}

	// Initial values for the first station on each line
	T1[0] = e[0] + a[0][0]
	T2[0] = e[1] + a[1][0]

	// Fill T1 and T2 arrays with minimum times using dynamic programming approach
	for i := 1; i < 4; i++ {
		T1[i] = min(T1[i-1]+a[0][i], T2[i-1]+a[0][i]+ct[1][i])
		T2[i] = min(T2[i-1]+a[1][i], T1[i-1]+a[1][i]+ct[0][i])
	}
	fmt.Println("After minimum T1 :", T1)
	fmt.Println("After minimum T2 : ", T2)

	// Calculate the minimum time to exit from both lines
	n := 4
	minimum := min(T1[n-1]+x[0], T2[n-1]+x[1])

	return minimum
}

func main() {
	// Given input data
	a := [2][4]int{{4, 5, 3, 2},
		{2, 10, 1, 4}}

	ct := [2][4]int{{0, 7, 4, 5},
		{0, 9, 2, 8}}

	e := [2]int{10, 12}

	x := [2]int{18, 7}

	fmt.Println("Assembly Line : ", a)
	fmt.Println("Extra time when change the path : ", ct)
	fmt.Println("Time taken to enter in the line : ", e)
	fmt.Println("Time taken to exit from the line : ", x)

	// Calculate the minimum distance using carAssemblyLineScheduling function
	var resp int
	resp = carAssemblyLineScheduling(a, ct, e, x)
	fmt.Println("Minimum distance is : ", resp)
}

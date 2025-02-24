package main

import "fmt"

func main() {
	var r, c int
	fmt.Println("enter number of rows and column for your matrix : ")
	fmt.Scan(&r, &c)

	matrix := make([][]int, r)

	for i := range matrix {
		matrix[i] = make([]int, c)
	}

	fmt.Println("enter all your matrix elements : ")
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	fmt.Println("Okay! so this is your matrix : ")
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

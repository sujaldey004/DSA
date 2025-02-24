package main

import "fmt"

func main() {
	var r, c int
	fmt.Println("Enter row and column : ")
	fmt.Scan(&r, &c)

	matrix := make([][]int, r)

	for i := range matrix {
		matrix[i] = make([]int, c)
	}

	fmt.Println("Enter matrix element : ")
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	sum := 0
	n := 0
	for i := 0; i < r; i++ {
		sum += matrix[i][n]
		n++
	}

	fmt.Println("Sum of Diagonal of matrix :", sum)
}

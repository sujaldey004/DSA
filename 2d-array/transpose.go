package main

import "fmt"

func main() {
	var r, c int
	fmt.Println("Enter rows and column : ")
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

	fmt.Println("this is the transpose matrix : ")
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			fmt.Print(matrix[j][i], " ")
		}
		fmt.Println()
	}
}

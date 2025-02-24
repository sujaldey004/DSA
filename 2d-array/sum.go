package main

import "fmt"

func main() {
	var r, c int
	fmt.Println("Enter rows and columns for your matrix : ")
	fmt.Scan(&r, &c)

	matrix := make([][]int, r)

	for i := range matrix {
		matrix[i] = make([]int, c)
	}

	fmt.Println("Enter your matrix element : ")
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	fmt.Println("This is your matrix : ")
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

	sum := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			sum += matrix[i][j]
		}
		fmt.Printf("Sum of row %d is : %d\n", i+1, sum)
		sum = 0
	}
}

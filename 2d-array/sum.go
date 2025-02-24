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

	for i := 0; i < c; i++ {
		sum := 0
		for j := 0; j < r; j++ {
			sum += matrix[j][i]
		}
		fmt.Printf("Sum of columns %d is : %d\n", i+1, sum)
		sum = 0
	}
}

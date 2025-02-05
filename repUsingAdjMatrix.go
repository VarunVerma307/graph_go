package main

import (
	"fmt"
)

func addedge(mat [][]int, i, j int) {
	mat[i][j] = 1
}

func display(mat [][]int) {
	V := len(mat)
	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println()
	}
}

// func main() {
// 	V := 4
// 	mat := make([][]int, V)
// 	for i := range mat {
// 		mat[i] = make([]int, V)
// 	}

// 	addedge(mat, 0, 1)
// 	addedge(mat, 0, 2)
// 	addedge(mat, 1, 2)
// 	addedge(mat, 2, 3)

// 	fmt.Println("Adjacency Matrix Representation")
// 	display(mat)
// }
package main

import (
	"fmt"
)


func addEdge(adj [][]int, i, j int) {
	adj[i] = append(adj[i], j)
}


func displayAdjList(adj [][]int) {
	for i := 0; i < len(adj); i++ {
		fmt.Printf("%d: ", i) 
		for _, j := range adj[i] {
			fmt.Printf("%d ", j) 
		}
		fmt.Println()
	}
}


// func main() {
// 	V:= 4
// 	adj := make([][]int, V)


// 	addEdge(adj, 0, 1)
// 	addEdge(adj, 0, 2)
// 	addEdge(adj, 1, 2)
// 	addEdge(adj, 2, 3)

// 	fmt.Println("Adjacency List Representation:")
// 	displayAdjList(adj)
// }

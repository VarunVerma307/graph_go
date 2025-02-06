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

func dfs(adj [][]int, i int, isvisited []bool, ans *[]int) {
	isvisited[i] = true
	for _, it := range adj[i] {
		if !isvisited[it] {
			dfs(adj, it, isvisited, ans)
		}
	}
	*ans = append(*ans, i)
}

func topologicalSort(adj [][]int, isvisited []bool, ans *[]int) {
	for i := 0; i < len(adj); i++ {
		if !isvisited[i] {
			dfs(adj, i, isvisited, ans)
		}
	}
}

func main() {
	V := 4
	adj := make([][]int, V)

	addEdge(adj, 0, 1)
	addEdge(adj, 0, 2)
	addEdge(adj, 1, 2)
	addEdge(adj, 2, 3)

	fmt.Println("Adjacency List Representation:")
	displayAdjList(adj)
	isvisited := make([]bool, V)
	var ans []int
	topologicalSort(adj, isvisited, &ans)

	fmt.Println("Topological Sort:")
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Printf("%d ", ans[i])
	}
}
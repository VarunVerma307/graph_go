package main

import (
	"fmt"
)

func dfsRec(adj [][]int, isvisited map[int]bool, src int) {
	isvisited[src] = true
	fmt.Print(src, " ")
	for _, i := range adj[src] {
		if isvisited[i] == false {
			dfsRec(adj, isvisited, i)
		}
	}
}
func dfs(adj [][]int, src int) {
	isvisited := make(map[int]bool)
	dfsRec(adj, isvisited, src)
}
func adedge(adj [][]int, i, j int) {
	adj[i] = append(adj[i], j)
}
func dispaly(adj [][]int) {
	for i := 0; i < len(adj); i++ {
		fmt.Printf("%d : ", i)
		for _, j := range adj[i] {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}

}

func main() {
	V := 5
	adj := make([][]int, V)
	adedge(adj, 1, 2)
	adedge(adj, 1, 0)
	adedge(adj, 2, 0)
	adedge(adj, 2, 3)
	adedge(adj, 2, 4)
	dispaly(adj)
	dfs(adj, 1)
}

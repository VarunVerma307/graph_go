package main

import (
	"fmt"
	"container/list"
)

func minSteps(m, n, d int) int {
	if d > max(m, n) {
		return -1
	}

	q := list.New()

	visited := make([][]bool, m+1)
	for i := range visited {
		visited[i] = make([]bool, n+1)
	}

	q.PushBack([]int{0, 0, 0})
	visited[0][0] = true

	for q.Len() > 0 {
		curr := q.Front()
		q.Remove(curr)

		state := curr.Value.([]int)
		jug1 := state[0]
		jug2 := state[1]
		steps := state[2]

		fmt.Printf("Step %d: Jug1 = %d, Jug2 = %d\n", steps, jug1, jug2)
		if jug2 == d {
			return steps
		}

		// 1: Fill jug1
		if !visited[m][jug2] {
			visited[m][jug2] = true
			q.PushBack([]int{m, jug2, steps + 1})
		}

		// 2: Fill jug2
		if !visited[jug1][n] {
			visited[jug1][n] = true
			q.PushBack([]int{jug1, n, steps + 1})
		}

		// 3: Empty jug1
		if !visited[0][jug2] {
			visited[0][jug2] = true
			q.PushBack([]int{0, jug2, steps + 1})
		}

		// 4: Empty jug2
		if !visited[jug1][0] {
			visited[jug1][0] = true
			q.PushBack([]int{jug1, 0, steps + 1})
		}

		// 5: Pour jug1 into jug2
		pour1to2 := min(jug1, n-jug2)
		if !visited[jug1-pour1to2][jug2+pour1to2] {
			visited[jug1-pour1to2][jug2+pour1to2] = true
			q.PushBack([]int{jug1 - pour1to2, jug2 + pour1to2, steps + 1})
		}

		// 6: Pour jug2 into jug1
		pour2to1 := min(jug2, m-jug1)
		if !visited[jug1+pour2to1][jug2-pour2to1] {
			visited[jug1+pour2to1][jug2-pour2to1] = true
			q.PushBack([]int{jug1 + pour2to1, jug2 - pour2to1, steps + 1})
		}
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var m, n, d int
	fmt.Println("Enter the capacity of smaller jug")
	fmt.Scan(&m)
	fmt.Println("Enter the capacity of larger jug")
	fmt.Scan(&n)
	fmt.Println("Enter the target capacity")
	fmt.Scan(&d)
	x := minSteps(m, n, d)
	fmt.Printf("Minimum number of steps are %d\n", x)
}
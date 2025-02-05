// #include <iostream>
// #include<vector>
// #include<queue>
// #include<map>
// using namespace std;

// void adedges(vector<vector<int>> &adj, int i,int j){
//     adj[i].push_back(j);
// }

// void display(vector<vector<int>>adj){
//     for(int i=0; i<adj.size(); i++){
//         cout<<i<<": ";
//         for(int j=0; j<adj[i].size(); j++){
//             cout<<j<<" ";
//         }
//         cout<<endl;
//     }
// }

// void bfs(vector<vector<int>>adj,int s){
//     queue<int> q;
//     vector<bool> m(adj.size(), false);
//     m[s]=true;
//     q.push(s);

//     while(!q.empty()){
//         int curr=q.front();
//         q.pop();
//         cout<<curr<<" ";

//          for (int x : adj[curr]) {
//             if (!m[x]) {
//                 m[x] = true;
//                 q.push(x);
//             }
//         }
//     }

// }

// int main(){
//     int V=4;
//     vector<vector<int>>adj (V);
//     adedges(adj,1,2);
//     adedges(adj,1,0);
//     adedges(adj,2,0);
//     adedges(adj,1,3);
//     adedges(adj,3,2);
//     adedges(adj,3,1);
//     display(adj);
//     bfs(adj,1);

// }

package main

import (
	"fmt"
)

type queue struct {
	List []int
}

func (q *queue) enque(element int) {
	q.List = append(q.List, element)
}

func (q *queue) deque() int {
	if q.empty() {
		fmt.Println("Queue is empty\n")
		return -1
	} else {
		element := q.List[0]
		q.List = q.List[1:]
		return element
	}
}

func (q *queue) empty() bool {
	return len(q.List) == 0
}

func bfs(adj [][]int, s int) {
	isvisited := make(map[int]bool)
	var bfsq queue

	isvisited[s] = true
	bfsq.enque(s)

	for !bfsq.empty() {
		curr := bfsq.List[0]
		fmt.Println(curr, " ")
		for _, v := range adj[curr] {
			if !isvisited[v] {
				bfsq.enque(v)
				isvisited[v] = true
			}
		}
		bfsq.deque()
	}
}

func adddge(adj [][]int, i, j int) {
	adj[i] = append(adj[i], j)
}

func displa(adj [][]int) {
	for i := 0; i < len(adj); i++ {
		fmt.Printf("%d: ", i)
		for _, j := range adj[i] {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}

func main() {
	V := 4
	adj := make([][]int, V)

	adddge(adj, 0, 1)
	adddge(adj, 0, 2)
	adddge(adj, 1, 2)
	adddge(adj, 2, 3)

	fmt.Println("Adjacency List Representation:")
	displa(adj)
	bfs(adj, 0)
}

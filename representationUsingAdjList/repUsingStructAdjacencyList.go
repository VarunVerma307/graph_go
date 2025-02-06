package main

import (
	"fmt"
)

type graph struct{
   vertices []*Vertex
}

type Vertex struct{
	key int 
	adjacents []*Vertex
}

func (g *graph) addVertex(k int){
	if contains(k,g.vertices){
		fmt.Println("Vertex can't be added as a vertex with the same value already exist")		
	}else{
		g.vertices = append(g.vertices, &Vertex{key: k})
	}

}

func contains (k int, s []*Vertex) bool{
    for _,v := range s{
		if v.key == k {
			return true
		}
	}
	return false
}

func (g *graph) addEdge(from ,to int){
	fromvertex := g.getVertex(from)
	tovertex := g.getVertex(to)
    if fromvertex==nil || tovertex==nil {
		fmt.Println("Cannot cretae this edge as one of the vertex does'nt exist")
		return
	}else if contains (to,fromvertex.adjacents ){
        fmt.Println("Edge already exists")
	}else{
		fromvertex.adjacents = append(fromvertex.adjacents,tovertex)
	}
	}
	
func (g *graph) getVertex(k int) *Vertex{
	for i,v := range g.vertices{
		if k == v.key {
			return g.vertices[i]
		}
	}
		return nil
}

func (g *graph) print(){
	for _,v := range g.vertices{
		fmt.Printf("\nVertex %v : ", v.key)
		for _,v := range v.adjacents{
			fmt.Printf("%v",v.key)
		}
	}
	fmt.Println()
}


// func main(){
// 	fmt.Println("hello world")
// 	test := &graph{}
// 	for i:=0 ; i<5 ; i++{
//          test.addVertex(i)
// 	}
// 	test.addEdge(1,5)
// 	test.addEdge(1,2)
// 	test.addEdge(1,3)
// 	test.addEdge(1,4)
// 	test.addEdge(1,3)
// 	test.print()
// }




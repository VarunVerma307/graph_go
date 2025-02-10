package main

import (
	"fmt"
)

func pour(fromcap, tocap, d int) int{
   from := fromcap
   to := 0
   step := 1
   fmt.Printf("Step %d: (%d,%d)\n", step, from, to)

   for from!=d && to!=d{
	  temp:= min(from,tocap-to)
	  from-=temp
	  to+=temp
	  step++;
      fmt.Printf("Step %d: (%d,%d)\n", step, from, to)

	  if from==d || to ==d {
		break
	  }

	  
	  if from==0 {
		from=fromcap
		step++
		fmt.Printf("Step %d: (%d,%d)\n", step, from, to)
	  }

	  if to==tocap{
		to=0
		step++
		fmt.Printf("Step %d: (%d,%d)\n", step, from, to)
	  }
   }
   return step
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func minsteps(m,n,d int) int{
	if m>n{
		n,m=m,n
	}
	if d>n{
		return -1
	}
	if d%gcd(m,n)!=0{
		return -1
	}
	return pour(m,n,d)
}

// func main() {
// 	var m, n, d int
// 	fmt.Println("Enter the capacity of first jug")
// 	fmt.Scan(&m)
// 	fmt.Println("Enter the capacity of second jug")
// 	fmt.Scan(&n)
// 	fmt.Println("Enter the target capacity")
// 	fmt.Scan(&d)
// 	var x int = minsteps(m, n, d)
// 	fmt.Println("\nMinimum steps are:", x)
// }
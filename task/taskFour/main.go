package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	first := n / 1000   
	last := n % 1000

	s1 := first/100 + (first/10)%10 + first%10

	s2 := last/100 + (last/10)%10 + last%10

	if (s1 == s2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
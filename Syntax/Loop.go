package main

import "fmt"

/*
	note: 在go中 沒有++i

*/
func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Printf("\n")

	for j := 0; j < 10; {
		j += 1
		fmt.Printf("%d", j)
	}
	fmt.Printf("\n")

}

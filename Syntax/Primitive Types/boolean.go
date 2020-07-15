package main

import "fmt"

func main() {
	var boolParam bool = true
	a := true
	b := false

	fmt.Println(boolParam)

	if boolParam {
		fmt.Println("test !")
	}

	fmt.Println(a || b)
}

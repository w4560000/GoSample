package main

import "fmt"

func main() {
	Test1("a", "b", "c")

	fmt.Println(Test2(), Test3())

	// 建立 local func  類似c# func
	customFunc := func(p1, p2 int) int {
		return p1 - p2
	}

	fmt.Println(customFunc(100, 50))
}

// 類似c# param
func Test1(key ...string) {
	// 型別是 array
	fmt.Printf("%T\n", key)
	fmt.Println(key)
}

// go的func 寫法 return 值在後面
func Test2() int {
	return 1
}

// return值可以指定變數
func Test3() (a int) {
	a = 87
	return
}

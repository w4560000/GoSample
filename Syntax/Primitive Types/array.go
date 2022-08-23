package main

import (
	"fmt"
)

func main() {
	func1()
	fmt.Println("-------------")
	func2()
}

func func1() {
	// 宣告陣列並賦值
	var arr [3]string
	arr[0] = "Iron Man"
	arr[1] = "Dr.Stange"
	arr[2] = "Thor"

	fmt.Print("arr =")
	fmt.Println(arr) // 印出: [Iron Man Dr.Stange Thor]

	// 初始化
	var arr1 = [3]string{"Iron Man", "Dr.Stange", "Thor"}
	fmt.Print("arr1 = ")
	fmt.Println(arr1) // 印出: [Iron Man Dr.Stange Thor]

	// 宣告陣列時若有設定初始值 則可省略陣列長度
	arr2 := [...]string{"Iron Man", "Dr.Stange", "Thor"}
	fmt.Print("arr2 = ")
	fmt.Println(arr2)
	fmt.Print("arr2 length = ")
	fmt.Println(len(arr2))
	// 印出: [Iron Man Dr.Stange Thor]
}

func func2() {
	arr := [...]string{"Iron Man", "Dr.Stange", "Thor"}

	// k為元素索引，v為元素值
	for k, v := range arr {
		fmt.Println(k, v)
	}
}

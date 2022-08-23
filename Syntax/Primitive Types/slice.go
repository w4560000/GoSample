package main

import (
	"fmt"
	"reflect"
)

func main() {
	// slice 基本宣告
	// func1()
	// fmt.Println("-------------")

	func2()

	// 測試 array、slice pass by value or pass by Ref
	//testPassbyValueOrPassbyRef()
}

func testArray(array [5]int) {
	fmt.Println("modifyArray1", array, fmt.Sprintf("%p", &array))
	array[3] = 3
	fmt.Println("modifyArray2", array, fmt.Sprintf("%p", &array))
}

func modifySlice(slice []int) {
	fmt.Println("modifySlice1", slice, fmt.Sprintf("%p", &slice))
	slice[3] = 3
	fmt.Println("modifySlice2", slice, fmt.Sprintf("%p", &slice))
}

func func1() {
	// 宣告字串陣列，並且初始化 5 個英雄名稱
	arr := [5]string{"Iron Man", "Dr.Stange", "Thor", "Captain America", "Hulk"}

	// 建立字串切片，並且取得arr陣列的第 1 個到第 3 個元素
	// [ 開始位置 : 結束位置 ]
	slice := arr[1:4]

	// 印出切片元素 k = index , v = value
	for k, v := range slice {
		fmt.Println(k, v)
	}

	// 修改 slice時 也會同時修改到arr，因為 slice 實際上 不是一份新的
	fmt.Println("----- 改為 Bat Man -----")
	slice[0] = "Bat Man"

	fmt.Println("原始arr")
	// 印出切片元素 k = index , v = value
	for k, v := range arr {
		fmt.Print(k, " ", v)
	}

	fmt.Println("slice")
	// 印出切片元素 k = index , v = value
	for k, v := range slice {
		fmt.Print(k, " ", v)
	}
}

func func2() {
	// 有指定陣列長度為 array (在此以...取代長度，go 會依照初始值數量建立相同長度)
	arr := [...]string{"Iron Man", "Dr.Stange", "Thor"}
	//arr = append(arr, "123") // array 因指定了固定陣列長度 而無法append

	// 印出切片元素 k = index , v = value
	for k, v := range arr {
		fmt.Print(k, " ", v)
	}
	fmt.Println("")
	// 無指定陣列長度為 slice
	slice := []string{"Iron Man", "Dr.Stange", "Thor"}
	slice = append(slice, "1") // slice 可自動增長長度

	// 印出切片元素 k = index , v = value
	for k, v := range slice {
		fmt.Print(k, " ", v)
	}
	fmt.Println("")
	fmt.Println(reflect.TypeOf(arr).Kind())
	fmt.Println(reflect.TypeOf(slice).Kind())
}

func testPassbyValueOrPassbyRef() {
	// 測試 array pass by value or pass by ref
	// 結果 pass by value
	a1 := [5]int{0, 0, 0, 0, 0}
	fmt.Println("array", a1, fmt.Sprintf("%p", &a1))
	testArray(a1)
	fmt.Println("array", a1, fmt.Sprintf("%p", &a1))
	// 當傳入 array 到別的fuc時，是pass by value (新的ref位址)
	// 所以在別的fuc有異動array時，不影響到原有的array，因為ref不同
	// 印出:
	// array [0 0 0 0 0] 0xc00000e3c0
	// modifyArray [0 0 0 3 0] 0xc00000e420
	// array [0 0 0 0 0] 0xc00000e3c0

	// 測試 slice pass by value or pass by ref
	// 結果 pass by value
	sa1 := [5]int{0, 0, 0, 0, 0}
	s1 := sa1[0:4]
	fmt.Println("slice", s1, fmt.Sprintf("%p", &s1))
	modifySlice(s1)
	fmt.Println("slice", s1, fmt.Sprintf("%p", &s1))
	fmt.Println("原始array", sa1, fmt.Sprintf("%p", &sa1))
	// slice 比較特別，slice 的資料 是指向原始array
	// 而儘管 slice 是 pass by value，傳到別的func後，就算 ref 被更改了，但實際上他還是指向原始array
	// 所以修改了slice 也同時會異動到原始array
	// 印出:
	// slice [0 0 0 0] 0xc000008078
	// modifySlice1 [0 0 0 0] 0xc0000080a8
	// modifySlice2 [0 0 0 3] 0xc0000080a8
	// slice [0 0 0 3] 0xc000008078
	// 原始array [0 0 0 3 0] 0xc00000e4e0
}

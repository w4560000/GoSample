package main

import (
	"fmt"

	"GOSAMPLE/Syntax/bxtest"
)

func main() {

	// 正常宣告
	var param1 string = "param1"

	// 不宣告型別 由編譯器動態判斷
	var param2 = "param2"

	// 宣告的語法糖
	param3 := "param3"

	// 以const宣告的變數為常數 無法變更
	const param4 string = "param4"

	// 一次定義多個變數
	var param5, param6 = "123", 987

	// 一次定義多個變數
	var (
		param7   = "param5"
		intParam = 123
	)

	//swap參數，傳統寫法要swap參數時還需要一個第三方暫存變數來暫存 在go就不用
	var a, b = 1, 2
	a, b = b, a

	// 接收func參數
	funcData1, funcData2 := returnData()

	// 當有參數不需要使用時 可使用匿名佔位符號 '_'  用匿名佔位符號宣告的變數 無法被使用
	funcData3, _ := returnData()

	fmt.Println(param1, param2, param3, param4, param5, param6, param7, intParam, bxtest.TestParam, a, b)

	fmt.Println(funcData1, funcData2, funcData3)
}

func returnData() (string, string) {
	return "data1", "data2"
}

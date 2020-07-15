package main

import "fmt"

/*  go的流程控制有兩種

if/else
case1中是常見的if/else 在go中不需要上下括號

case2中可以if條件加上一行表達式

case3中跟其他語言一樣可以在條件裡宣告變數

switch case
case1中也是常見的switch case用法 但在go中不用再寫break

case2中比較特殊 以C#的switch case來說 不能用動態判斷式去判別 而在go中可以 方便很多

case3中在case中可以允許多個條件 在c#8.0中也有實現

*/
func main() {
	myAge := 20

	// if/else 1
	if myAge < 13 {
		fmt.Println("if/else 1 Child")
	} else if myAge >= 13 && myAge < 20 {
		fmt.Println("if/else 1 Teen")
	} else {
		fmt.Println("if/else 1 Adult")
	}

	// if/else 2
	if myAge = myAge - 10; myAge < 13 {
		fmt.Println("if/else 2 Child")
	} else if myAge >= 13 && myAge < 20 {
		fmt.Println("if/else 2 Teen")
	} else {
		fmt.Println("if/else 2 Adult")
	}

	// if/else 3
	if testAge := 60; testAge < 13 {
		fmt.Println("if/else 3 Child")
	} else if testAge > 50 {
		fmt.Println("if/else 3 OldMan")
	}

	// switch case 1
	testParam := 3
	switch testParam {
	case 1:
		fmt.Println("switch case 1 First")
	case 2:
		fmt.Println("switch case 1 Second")
	case 3:
		fmt.Println("switch case 1 Third")
	}

	// switch case 2
	mySonAge := 5
	switch {
	case mySonAge < 13:
		fmt.Println("switch case 2 Child")
	case mySonAge >= 13 && mySonAge < 20:
		fmt.Println("switch case 2 Teen")
	case mySonAge > 20:
		fmt.Println("switch case 2 Adult")
	}

	// switch case 3
	flag := 30
	switch flag {
	case 0, 10:
		fmt.Println("switch case 3 Zero - Ten")
	case 20, 30:
		fmt.Println("switch case 3 Twenty - Thirty")

	}

}

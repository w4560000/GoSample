package bxtest

// 當要給外部引用時 首位字母必須為大寫才可以被外部引用
// TestParam 可以被引用 而 testParam 不行
var TestParam = "Test"
var testParam = "test1"

// func也是一樣 首字小寫無法被引用
func Add(a, b int) int {
	return a + b
}

func addtest(a, b int) int {
	return a + b
}

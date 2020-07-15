package main

import (
	"fmt"
	"strings"
)

func main() {
	testString1 := "123"

	// 定義多行文字 類似 c# 宣告字串前的@
	testString2 := `123
	ss
	ass
	s`
	fmt.Println(testString1)
	fmt.Println(testString2)

	// string extension
	fmt.Println(
		// 是否包含
		strings.Contains("test", "st"),

		// stringA 有幾個字元與 StringB 相符
		strings.Count("test", "t"),

		// StringA 的前綴字 是否為 StringB
		strings.HasPrefix("test", "te"),

		// StringA 的後綴字 是否為 StringB
		strings.HasSuffix("test", "st"),

		// 查詢 StringB 在 StringA中的index
		strings.Index("test", "e"),

		// 將string array 以 '-'組合成字串
		strings.Join([]string{"a", "b"}, "-"),

		// 重複建立 StringA 5次 並串接起來
		strings.Repeat("a1", 5),

		// 將StringA 中的 'a' 以 'b'代替 ， 替換次數為1次
		strings.Replace("aaaa", "a", "b", 1),

		// 將String A 以 '-' 切開來， 回傳型別為 []string
		strings.Split("a-b-c-d-e", "-"),

		// 將String A 轉為小寫
		strings.ToLower("TEST"),

		// 將String A 轉為大寫
		strings.ToUpper("test"),
	)
}

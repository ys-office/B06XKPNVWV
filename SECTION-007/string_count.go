package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// string 型の変数を宣言
	var ja string = "Go言語"

	// 文字列の長さを出力
	fmt.Println(ja, " len:", utf8.RuneCountInString(ja))
}

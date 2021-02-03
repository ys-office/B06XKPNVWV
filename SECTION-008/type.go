package main

import "fmt"

func main() {

	// int 型から myInteger 型を宣言
	type myInteger int

	// myInteger 型の変数を作成
	var i myInteger = 123

	// int 型と同じように演算が可能
	i += 1

	// 出力
	fmt.Println(i)

	// 新しい構造体を作成
	type myStruct struct {
		a int
		b int
	}
}

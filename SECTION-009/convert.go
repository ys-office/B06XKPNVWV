package main

import "fmt"

func main() {

	// int 型の変数を宣言
	var i int = 1234

	// int 型から uint32 型への変換
	var u uint32 = uint32(i)

	// uint 型から float32 型への変換
	var f float32 = float32(u)

	// int 型から string 型への変換
	var s string = string(i)

	// string 型から配列（正確にはスライス）型への変換
	var b []byte = []byte("abc")

	// 出力
	fmt.Println(u)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(b)
}

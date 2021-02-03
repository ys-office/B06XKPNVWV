package main

import "fmt"

func main() {

	// int 型の変数を宣言
	var i int = 12345

	// int 型から int64 への変換
	var i64 int64 = int64(i)

	// 出力
	fmt.Println(i64)
}

package main

import "fmt"

func main() {

	// bool 型の変数を宣言
	var b bool

	// bool 型の変数にリテラル定数 true と false を代入
	b = true
	b = false

	// bool 型の変数に論理演算した結果を代入
	b = true || false

	// 出力
	fmt.Println(b)
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("型指定と出力、型確認")
	var num int = 42
	fmt.Println(num)
	fmt.Println(reflect.TypeOf(num))

	// 改行あり
	fmt.Println() //改行
	fmt.Println("改行あり")
	fmt.Println(42)
	fmt.Println(42)

	// 改行なし
	fmt.Println() //改行
	fmt.Println("改行なし")
	fmt.Print(42)
	fmt.Print(42)

	// フォーマット済み文字列
	fmt.Println("\n") //改行
	fmt.Println("フォーマット済み文字列")
	var age int = 42
	fmt.Printf("父は%d歳である\n", age)

	// キーボードからの入力を取得
	fmt.Println() //改行
	var key int
	fmt.Scan(&key)
	fmt.Println(key)
}

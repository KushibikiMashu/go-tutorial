package main

import "fmt"

func main() {
	// 空のインターフェースは、任意の型の値を保持できます。
	// (全ての型は、少なくともゼロ個のメソッドを実装しています。)
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

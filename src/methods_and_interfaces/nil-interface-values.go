package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// nil インターフェースの値は、値も具体的な型も保持しません。
	// nil インターフェースのメソッドを呼び出すと、ランタイムエラーになります。
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// /go/src/methods_and_interfaces # go run nil-interface-val
// ues.go
// (<nil>, <nil>)
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x488ac1]

// goroutine 1 [running]:
// main.main()
//         /go/src/methods_and_interfaces/nil-interface-values.go:14 +0x91
// exit status 2

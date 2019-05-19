package main

import (
	"fmt"
	"math"
)

// 例で挙げたstructの型だけではなく、任意の型(type)にもメソッドを宣言できます。
type MyFloat float64

// レシーバを伴うメソッドの宣言は、レシーバ型が同じパッケージにある必要があります。
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main()  {
	f := MyFloat(-math.Sqrt)
	fmt.Println(f.Abs())
}

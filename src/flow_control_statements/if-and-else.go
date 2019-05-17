package main

import (
	"fmt"
	"math"
)

// if ステートメントで宣言された変数は、 else ブロック内でも使うことができます。
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main()  {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

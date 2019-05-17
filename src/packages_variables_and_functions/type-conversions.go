package main

import (
	"fmt"
	"math"
)

// C言語とは異なり、Goでの型変換は明示的な変換が必要です
func main()  {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x,y,z)
}
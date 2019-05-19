package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main()  {
	// interface(インタフェース)型は、メソッドのシグニチャの集まりで定義します。
	// そのメソッドの集まりを実装した値を、interface型の変数へ持たせることができます。
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	fmt.Println(a.Abs())
	
	a = &v // a *Vertex implements Abser

	// a Vertex does not implement Abser
	// a = v 

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 9 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

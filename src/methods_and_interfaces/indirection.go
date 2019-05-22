package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// メソッドでポインタレシーバが自動的に呼びだされます。
	// Scale メソッドはポインタレシーバを持つ場合、Goは利便性のため、
	// v.Scale(5) のステートメントを (&v).Scale(5) として解釈します。
	v.Scale(2)
	ScaleFunc(&v, 10)

	// メソッドがポインタレシーバである場合、
	// 呼び出し時に、変数、または、ポインタのいずれかのレシーバとして取ることができます
	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// スライスは配列への参照のようなものです。
// スライスはどんなデータも格納しておらず、単に元の配列の部分列を指し示しています。
// スライスの要素を変更すると、その元となる配列の対応する要素が変更されます。
// 同じ元となる配列を共有している他のスライスは、それらの変更が反映されます。

// /go/src/structs_slices_and_maps # go run slices-pointers.go
// [John Paul George Ringo]
// [John XXX] [XXX George]
// [John XXX George Ringo]

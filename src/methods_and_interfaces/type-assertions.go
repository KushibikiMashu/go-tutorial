package main

import "fmt"

func main()  {
	// 型アサーション は、インターフェースの値の基になる具体的な値を利用する手段を提供します。
	var i interface{} = "hello"

	s, ok := i.(string)
	fmt.Println(s, ok)
	
	t, ok := i.(string)
	fmt.Println(t, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// panic
	// f = i.(float64) 
	// fmt.Println(f)	
}

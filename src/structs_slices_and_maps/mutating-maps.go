package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// もし、 m に key があれば、変数 ok は true となり、存在しなければ、 ok は false となります。
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m["Answer"] = 40
	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}

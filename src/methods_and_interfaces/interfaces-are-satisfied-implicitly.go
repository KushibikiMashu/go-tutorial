package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

//  インタフェースを実装することを明示的に宣言する必要はありません
// ( "implements" キーワードは必要ありません)。
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}

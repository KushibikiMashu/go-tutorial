package main

import "fmt"

func sum (s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main()  {
	s := []int{7, 2, 8, -9, 4, 0}

	// マップとスライスのように、チャネルは使う前に以下のように生成します
	c := make(chan int)
	// 通常、片方が準備できるまで送受信はブロックされます。これにより、明確なロックや条件変数がなくても、goroutineの同期を可能にします。
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

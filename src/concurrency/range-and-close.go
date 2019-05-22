package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 送り手は、これ以上の送信する値がないことを示すため、チャネルを close できます
	// 注意: 送り手のチャネルだけをcloseしてください。受け手はcloseしてはいけません。 もしcloseしたチャネルへ送信すると、パニック( panic )します
	// チャネルは、ファイルとは異なり、通常は、closeする必要はありません。 closeするのは、これ以上値が来ないことを受け手が知る必要があるときにだけです
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// ループの for i := range c は、チャネルが閉じられるまで、チャネルから値を繰り返し受信し続けます
	for i := range c {
		fmt.Println(i)
	}
}

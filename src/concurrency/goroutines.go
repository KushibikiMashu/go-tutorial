package main

import (
	"fmt"
	"time"
)

// goroutine (ゴルーチン)は、Goのランタイムに管理される軽量なスレッドです。
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	go say("!")
	say("hello")
}

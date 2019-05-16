package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Hello, World")
	
	// Playground上では"2009-11-10 23:00:00 UTC"が出力される
	fmt.Println("The time is", time.Now())
}
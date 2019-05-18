package main

import "fmt"

func main()  {
	defer fmt.Println("mine")
	defer fmt.Println("is")

	fmt.Println("world")
}

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main()  {
	// 関数は error 変数を返します。そして、呼び出し元はエラーが nil かどうかを確認することでエラーをハンドル(取り扱い)します。
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"runtime"
)

// Go の switch は C や C++、Java、JavaScript、PHP の switch と似ていますが、
//  Go では選択された case だけを実行してそれに続く全ての case は実行されません。
// これらの言語の各 case の最後に必要な break ステートメントが Go では自動的に提供されます
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("%s", os)
	}
}

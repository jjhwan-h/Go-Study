package main

import (
	"fmt"
	"runtime"
)

func main() {
	num := 2
	if num < 1 { //if 초기문; 조건문
		fmt.Println("hello")
	} else if num >= 1 && num < 3 {
		fmt.Println("bye")
	} else {
		fmt.Println("hi")
	}

	switch os := runtime.GOOS; os { //swtich 초기문; 비교값
	case "darwin": //case "darwin", "linux"처럼 여러개도 가능
		fmt.Println("OS X.") //자동으로 break됨
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s. \n", os)
	}
}

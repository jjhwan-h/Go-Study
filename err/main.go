package main

import (
	"fmt"
)

//	func divide(x, y int) int {
//		return x / y
//	}
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("Error: divide by zero")
	} else {
		return x / y, nil
	}
}

func main() {
	// val, err := divide(10, 0) //err발생
	// if err != nil {
	// 	panic("Error: 0으로 나누었어요") //panic이 발생하면 프로그램이 강제종료
	// } else {
	// 	fmt.Println(val, err)
	// }

	defer func() {
		//만약 panic이 발생했을때 프로그램을 중단하지않고 계속 진행하고 싶다면
		// defer와 recover를 사용하여 코드를 계속이어나간다..
		r := recover()
		fmt.Println(r)
		fmt.Println("recored from panic")
	}()

	panic("panic: Something went wrong")

	//fmt.Println(errors.New("error Message"))
	// fmt.Println(fmt.Errorf("error Message"))

}

package main

import "fmt"

func main() {
	var num int
	num = add(1, 2)
	fmt.Println(num)

	var num2, num3 int
	num2, num3 = calc(5, 10) //반환값이 두개
	fmt.Println(num2, num3)

	num4 := add2(5, 6)
	fmt.Println(num4)
}
func add(x int, y int) int {
	return x + y
}

func calc(x int, y int) (int, int) { //반환값 두개도 가능
	return x + y, x - y
}

func add2(x, y int) (val int) { //반환할 변수를 지정
	val = x + y
	return //val가 자동으로 return
}

func addOne(num *int) {
	*num += 1 //포인터가 가리키는 전달받은 변수 참조하여 +1시킨다.
}

func addOne2(num int) {
	num += 1 //해당 함수의 scope내에서의 num
}

/*
func 함수명(매개변수 매변수타입)리턴타입{
	실행코드
	return 리턴값
}
*/

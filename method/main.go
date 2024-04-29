package main

import "fmt"

type Calculator struct {
	X int
}

func (c Calculator) add(y int) int { //리시버에 지정된 struct에 대한 메서드
	return c.X + y
}

func (c *Calculator) add2(y int) { //리시버를 포인터로 지정
	c.X += y
}

func main() {
	calc := Calculator{10} //중괄호로 구조체 초기화 및 선언
	val := calc.add(50)    //calc의 X값에 50을 더하여 반환

	calc2 := Calculator{X: 10}
	calc2.add2(20) //calc2의 X값을 참조하여 +20

	fmt.Println(val)
	fmt.Println(calc2.X)

}

/*
func (리시버) 메서드명(매개변수 매개변수 타입)변환타입{ //리시버는 구조체를 지정.
		실행코드
		return 반환값
	}
*/

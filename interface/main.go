package main

import "fmt"

/*
interface는 구현한 메서드를 묶어둔것
interface를 사용할 구조체는 interface를 구현해야하는데
go lang에서는 명시적으로 Interface를 지원하지 않는다.
interface내 메서드를 모두 구현하면 interface를 구현한것으로 간주한다. => 덕타이핑
*/

type Bird struct {
	name string
}

type Ducky interface {
	DuckSound() string //반환값 string
	DuckWalk() string
	isSwim() string
}

func (b Bird) DuckSound() string {
	return "꽥"
}
func (b Bird) DuckWalk() string {
	return "뒤뚱"
}
func (b Bird) isSwim() string {
	if b.name == "오리" {
		return "있다"
	} else {
		return "없다"
	}
}

func typeCheck(i interface{}) { //blank한 interface를 사용하면 어떤 값이든 받을 수 있다.
	fmt.Printf("%T\n", i)
}
func main() {
	Duck := Bird{name: "오리"}
	fmt.Println(Duck.DuckSound())
	fmt.Println(Duck.DuckWalk())
	fmt.Println(Duck.isSwim())

	var num1 int8 = 1
	var num2 int16 = 2
	var num3 int32 = 3
	var num4 int64 = 4
	var num5 uint8 = 5
	typeCheck(num1)
	typeCheck(num2)
	typeCheck(num3)
	typeCheck(num4)
	typeCheck(num5)

}

/*
type 인터페이스명 interface{
	메서드 집합
}
*/

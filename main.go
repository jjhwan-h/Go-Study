package main

import (
	"fmt" //input, output
)

func main() {

	// /*변수와 상수*/
	// var num int = 10
	// var num1, num2 int = 10, 20
	// //num1, num2 := 10, 20  //왈루스 기호
	// var (
	// 	num3 int = 10
	// 	num4 int = 20
	// )

	// const (
	// 	a = iota //0부터 순차적으로 1씩 증가
	// 	b
	// 	c
	// 	d
	// )
	// fmt.Println(a, b, c, d)
	// fmt.Println(num)

	// /*타입*/
	// var isBool bool = false
	// //var num5 int8 = 129 => 127을 넘어간 숫자를 저장 하므로 error
	// var num6 uint = 129
	// var float1 float64 = 20.2 //float64 64비트를 할당한다는 뜻
	// var byte1 byte            //1byte저장 0~255 uint8과 동일
	// var str1 string = "hello world"

	// //직접 type 지정
	// type MyInt int
	// var num MyInt = 10

	// /* 형변환 */

	// var float2 float32 = 3.14
	// num7 := int(float2)
	// fmt.Println("%T \n", num7) //type확인
	// fmt.Println(num7)          //3

	// data := strconv.FormatInt(int64(num7), 2) //부호있는 정수를 2진수 문자열로 변환

	// /*배열*/
	// var arr [3]int
	// //arr := {1,2,3}

	// arr = [3]int{1, 2, 3} //1,2,3을 대입
	// fmt.Println(arr)
	// fmt.Println("%T \n", arr)

	/*슬라이스*/
	//동적 배열느낌

	var s []int = []int{1, 2, 3}
	//s := []int{1,2,3,4,5}
	//s1 := make([]int,10) //10개의 원소 0으로 초기화
	fmt.Printf("%T \n", s) //[]int 슬라이스 자료형

	s = make([]int, 5) //크기를 변경하여 재선언
	fmt.Println(s)     //[0,0,0,0,0]

	s = []int{1, 2, 3, 4, 5}

	s[4] = 1
	s = append(s, 7) //7을 뒤에 push한다
	fmt.Println(s)   // 1,2,3,4,1,7

	fmt.Println(len(s)) //길이확인
	fmt.Println(cap(s)) //할당된 메모리확인

	s1 := make([]int, 5, 10) // 5만큼 0으로 초기화하고 사용, 10만큼의 배열할당, 10를 넘어서면 10만큼 추가할당
	fmt.Println(s1)          //[0,0,0,0,0]
	fmt.Println(cap(s1))     //10
	fmt.Println(len(s1))     //5

	s1 = append(s1, 11, 12, 13, 14, 15, 16, 17, 18)
	fmt.Println(s1)      //[0,0,0,0,0,11,12,13,14,15,16,17,18]
	fmt.Println(cap(s1)) //20
	fmt.Println(len(s1)) //13

	sl1 := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5)

	//copy로 슬라이스값 복사
	copy(sl2, sl1) //sl2에 sl1복사
	sl2[0] = 999
	fmt.Println(sl1)
	fmt.Println(sl2)
	//대입연산자로 복사시 참조값으로 복사(얕은복사)

	//슬라이싱

	sl3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(sl3[1:4]) //2,3,4

}

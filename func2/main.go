package main

import "fmt"

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}
func a() {
	i := 0
	defer fmt.Println(i) //0이 출력된다. defer가 선언될때의 값이 사용된다.
	i++

}
func c() (i int) {
	defer func() { i++ }() //익명함수 선언하고 바로실행
	return 1               //return 1이 실행되고 i는 1이된다.
	//defer의 익명함수가 실행되어 i는 2가되고 리턴된다.
}
func numbers(nums ...int) { //nums는 int를 여러개 가진다. 가변인자
	fmt.Println(nums)
	fmt.Printf("%T\n", nums)
}

func main() { //defer 키워드를 사용하면 함수가 종료되기 직전에 실행된다.
	defer fmt.Println("마지막으로 실행")
	defer fmt.Println("세번째로 실행")
	defer fmt.Println("두번째로 실행")
	fmt.Println("첫 번째로 실행")

	//a() //defer로 a()함수가 종료될때 0이 출력되게 된다.
	//b()

	/*
				a()와 b()의 실행결과
				첫 번째로 실행
		0
		3
		2
		1
		0
		두번째로 실행
		세번째로 실행
		마지막으로 실행
	*/
	fmt.Println(c())
	numbers(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	f := func() { //함수 리터럴 (익명함수)
		fmt.Println("Hello")
		//익명함수는 외부의 값을 참조값으로 가져온다.
	}
	f()
}

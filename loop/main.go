package main

import "fmt"

/*
비트코인의 스크립트는 튜링불완전
이더리움의 스마트컨트랙트는 튜링완전하다.
튜링완전의 조건 => 이론적으로 계산가능한 모든것을 할 수 있다.
*/

func main() {
	sum := 0

	for i := 0; i < 10; i++ { //초기문; 조건문; 사후구문 => 초기문 사후구문은 생략가능
		sum += i
	}
	fmt.Println(sum)

	i := 1
	for { //break를 적지않으면 무한반복
		if i > 10 {
			break
		} else {
			fmt.Println(i)
			i++
		}
	}
	i = 1
	for i < 10 {
		i++
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}

	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr { //range는 배열이나 슬라이스문 순회할때 사용
		fmt.Printf("indx:%d, value: %d\n", i, v)
	}
}

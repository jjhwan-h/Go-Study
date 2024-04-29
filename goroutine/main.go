package main

/*
go runtime에 의해 관리되는 경량 쓰레드
go routine을 사용하면 다수의 작업을 동시에 실행할 수 잇다.
*/
import (
	"fmt"
	"sync"
)

func printHorizon(wg *sync.WaitGroup) {
	defer wg.Done() //함수끝날 시 goroutine에서 제거
	for i := 0; i < 300; i++ {
		fmt.Print("==")
	}
}
func printVertical(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 300; i++ {
		fmt.Print("||")
	}
}

func checkOne(num int, c chan bool) {
	if num == 1 {
		c <- true
	} else {
		c <- false
	}
}

func main() {
	//printHorizon()
	//printVertical()
	var wg sync.WaitGroup

	// wg.Add(2) //goroutine2개 추가
	// go printHorizon(&wg)
	// go printVertical(&wg)

	num := 0
	ch := make(chan bool)

	fmt.Print("Enter a number:")
	fmt.Scan(&num)
	go checkOne(num, ch) //goroutine은 리턴값을 사용할 수 없다. goroutine이 끝날때까지 기다리고 값을 받는다. 별도의 wait필요x

	res := <-ch //channel은 큐구조롤 이루어져있다.
	fmt.Println(res)
	wg.Wait() //goroutine끝날때까지 기다린다.
}

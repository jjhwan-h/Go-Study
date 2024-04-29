package main

import "fmt"

func main() {
	num := 1

	var numPointer *int = &num
	//numPointer :=&num
	fmt.Println(&num)        //0x14000...
	fmt.Println(numPointer)  //0x14000...
	fmt.Println(*numPointer) //1
	*numPointer = 2
	fmt.Println(*numPointer)
	fmt.Println(num)
}

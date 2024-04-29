package main

import "fmt"

//input, output

func main() {
	var numMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}
	fmt.Println(numMap["one"]) //1

	// var numMap2 = make(map[string]int) //map[키]값
	// numMap2["one"] = 1

	delete(numMap, "one")
	fmt.Println(numMap["one"]) //0

	key, value := numMap["two"]

	fmt.Println(key, value) //2 true(value가 있는지 확인)
	key2, vlaue2 := numMap["one"]
	fmt.Println(key2, vlaue2) //0 false
}

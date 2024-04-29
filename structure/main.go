package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Vertex struct {
		X int
		Y int
	}

	var v0 Vertex = Vertex{1, 2}
	fmt.Println(v0)
	v1 := Vertex{1, 2}
	fmt.Println(v1)
	v2 := Vertex{Y: 2, X: 1}
	fmt.Println(v2)

	v2.Y = 10
	fmt.Println(v2)

	type Vertex2 struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	v3 := Vertex2{1, 2}
	data, err := json.Marshal(v3) //json 형태로 변환 data는 []byte err는 error
	fmt.Println(string(data), err)
}

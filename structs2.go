package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2} //Vertex 类型
	v2 = Vertex{X:1}  //Y:0 隐式
	v3 = Vertex{}     //X:0, Y:0 
	p =&Vertex{1, 2}  //*Vertex 类型
)

func main() {
	fmt.Println(v1, v2, v3, p)
	}
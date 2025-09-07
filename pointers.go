package main

import "fmt"

func main(){
	i,j := 42,2701

	p :=&i  //指向i的指针
	fmt.Println(*p) //通过指针读取 i
	*p = 21 //通过指针设置 i
	fmt.Println(i) //查看 i 的新值
	p = &j //指向j的指针
	*p = *p /37 //通过指针除以 j
	fmt.Println(j) //查看 j 的新值
}
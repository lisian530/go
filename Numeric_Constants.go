package main

import "fmt"

const (
	//通过将 1 位左移 100 位来创建一个巨大的数字。
	Big = 1 << 100
	// 再次向右移动 99 位，因此我们得到 1<<1，即 2。
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) 
package main

import (
	"fmt"
	hq4x "github.com/donatj/hqxgo"
)

func main() {
	x := make([]int, 2*2*4*4)

	hq4x.Hq4x_32_rb([]int{0, 1010, 2, 30}, x, 2, 2, 48, 7, 6, 0, false, false)

	fmt.Println(x)
}

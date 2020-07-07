package main

import (
	"fmt"
	"github.com/imroc/biu"
)

func main() {
	var a uint8
	a = 3
	// uint32 把最后一位取出来
	p := uint8(7)

	fmt.Println((a&1)<<p, biu.ToBinaryString((a & 1)), biu.ToBinaryString((a&1)<<p))

}

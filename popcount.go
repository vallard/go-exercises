package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

func init() {
	for i := range pc {
		fmt.Println(pc[i/2])
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	v := os.Args[1]
	i, _ := strconv.Atoi(v)
	n := PopCount(uint64(i))
	fmt.Println(n)
}

func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

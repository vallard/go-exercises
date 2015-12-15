package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", ""
	// Exercise 1.1
	start := time.Now()
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.2dns elapsed\n", time.Since(start).Nanoseconds())

	// Exercise 1.2
	s, sep = "", ""
	start = time.Now()
	for i := 1; i < len(os.Args[1:]); i++ {
		s += fmt.Sprintf("%v [%v] %s", sep, i, os.Args[i])
	}
	fmt.Println(s)
	fmt.Printf("%.2dns elapsed\n", time.Since(start).Nanoseconds())

	// Exercise 1.3
	s, sep = "", ""
	start = time.Now()
	s = strings.Join(os.Args[1:], " ")
	fmt.Println(s)
	fmt.Printf("%.2dns elapsed\n", time.Since(start).Nanoseconds())
}

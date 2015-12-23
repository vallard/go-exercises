package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func main() {
	var c Celsius = 100
	fmt.Printf("%g\n", c)
}

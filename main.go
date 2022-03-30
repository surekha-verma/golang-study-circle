package main

import (
	"fmt"
	"math"
)

type myFloat float64

func (m myFloat) ceil() float64 {
	return math.Ceil(float64(m))
}

func main() {
	num := myFloat(1.4)
	fmt.Println(num.ceil())
}

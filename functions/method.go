package function

/**
https://go.dev/tour/methods/1

Methods
Go does not have classes. However, you can define methods on types.
A method is a function with a special receiver argument.

The parenthesis before the function name is the Go way of defining the object on which these functions will operate
**/

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//If you want to modify the receiver use a pointer like:
func (v *Vertex) AbsPointer() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//If you dont need to modify the receiver you can define the receiver as a value like:
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func PrintMethod() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

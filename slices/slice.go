package slice

import "fmt"

func GetSlices() {
	var cities []string
	cities = make([]string, 20)
	fmt.Println(len(cities))

	// Declare a variable called i which is a slice of 5 int.
	i := make([]int, 20)
	// Declare a variable called f which is a slice of 9 float64.
	f := make([]string, 9)
	// Declare a variable called s which is a slice of 4 string.
	s := []int {1,2,3,4,5,6,7,7,8}

	fmt.Println(len(i), len(f), len(s))
}
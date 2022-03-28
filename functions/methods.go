package function

import (
	"fmt"
)

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

func PrintMethods() {
	animals := make([]Dog, 2)
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

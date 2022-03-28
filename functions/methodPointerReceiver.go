package function

// Go program to illustrate pointer receiver

import "fmt"

// authorStruct structure
type authorStruct struct {
	name      string
	branch    string
	particles int
}

// Method with a receiver of authorStruct type
func (a *authorStruct) show(abranch string) {
	(*a).branch = abranch
}

// Main function
func PrintMethodPointerReceiver() {

	// Initializing the values
	// of the author structure
	res := authorStruct{
		name:   "Sona",
		branch: "CSE",
	}

	fmt.Println("Author's name: ", res.name)
	fmt.Println("Branch Name(Before): ", res.branch)

	// Creating a pointer
	p := &res

	// Calling the show method
	p.show("ECE")
	fmt.Println("Author's name: ", res.name)
	fmt.Println("Branch Name(After): ", res.branch)
}

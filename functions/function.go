package function

/**
Function as a Field in Golang Structure
**/

// Finalsalary of function type
type Finalsalary func(int, int) int

// Creating structure
type Author struct {
	name      string
	language  string
	Marticles int
	Pay       int

	// Function as a field
	salary Finalsalary
}

// Creating structure
type AuthorWithAnonymousFunc struct {
	name      string
	language  string
	Tarticles int
	Particles int
	Pending   func(int, int) int
}

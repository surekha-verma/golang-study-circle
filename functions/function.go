package function

/**
type function_name func()
type strcut_name struct{
  var_name  function_name
}
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

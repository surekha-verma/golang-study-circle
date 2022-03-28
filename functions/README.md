# Functions and Methods

## Difference between Functions and Methods

* Go language support methods.
* Go methods are similar to Go function with one difference, i.e, the method contains a receiver argument in it

Difference Between Method and Function
| Method | Function |
| --- | --- |
| It contains a receiver.	 | It does not contain a receiver. |
| Methods of the same name but different types can be defined in the program. | Functions of the same name but different type are not allowed to be defined in the program. |
| It cannot be used as a first-order object.	 | It can be used as first-order objects and can be passed |


## Functions
Function is a collection of statements that perform some specific task and return the result to the caller. A function can also perform some specific task without returning anything. Function declaration means a way to construct a function.

    ```console
    func function_name(Parameter-list)(Return_type){
        // function body.....
    }
    ```

* Function as a Field in Golang Structure

    ```console
    type function_name func()
    type strcut_name struct{
    var_name  function_name
    }
    ```

    ```console
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
    ```

## Methods

Go does not have classes. However, you can define methods on types.
A method is a function with a special receiver argument.

The parenthesis before the function name is the Go way of defining the object on which these functions will operate
With the help of the receiver argument, the method can access the properties of the receiver. 

Here, the receiver can be of struct type or non-struct type.

### struct type receiver

    ```console
    type Animal interface {
        Speak() string
    }

    type Dog struct {
    }

    func (d Dog) Speak() string {
        return "Woof!"
    }

    type Cat struct {
    }

    func (c Cat) Speak() string {
        return "Meow!"
    }

    type Llama struct {
    }

    func (l Llama) Speak() string {
        return "?????"
    }

    type JavaProgrammer struct {
    }

    func (j JavaProgrammer) Speak() string {
        return "Design patterns!"
    }

    func main() {
        animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
        for _, animal := range animals {
            fmt.Println(animal.Speak())
        }
    }
    ```

[Reference](https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)
### Methods with Pointer Receiver

In Go language, you are allowed to create a method with a pointer receiver. With the help of a pointer receiver.

    ```console
    // Go program to illustrate pointer receiver
    package main

    import "fmt"

    // Author structure
    type author struct {
        name	 string
        branch string
        particles int
    }

    // Method with a receiver of author type
    func (a *author) show(abranch string) {
        (*a).branch = abranch
    }

    // Main function
    func main() {

        // Initializing the values
        // of the author structure
        res := author{
            name: "Sona",
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
    ```

** Output **
    ```console
    Author's name:  Sona
    Branch Name(Before):  CSE
    Author's name:  Sona
    Branch Name(After):  ECE
    ```

### Method Can Accept both Pointer and Value
As we know that in Go, when a function has a value argument, then it will only accept the values of the parameter, and if you try to pass a pointer to a value function, then it will not accept and vice versa. But a Go method can accept both value and pointer, whether it is defined with pointer or value receiver.

    ```console
    // Go program to illustrate how the
    // method can accept pointer and value

    package main

    import "fmt"

    // Author structure
    type author struct {
        name string
        branch string
    }

    // Method with a pointer
    // receiver of author type
    func (a *author) show_1(abranch string) {
        (*a).branch = abranch
    }

    // Method with a value
    // receiver of author type
    func (a author) show_2() {

        a.name = "Gourav"
        fmt.Println("Author's name(Before) : ", a.name)
    }

    // Main function
    func main() {

        // Initializing the values
        // of the author structure
        res := author{
            name: "Sona",
            branch: "CSE",
        }

        fmt.Println("Branch Name(Before): ", res.branch)

        // Calling the show_1 method
        // (pointer method) with value
        res.show_1("ECE")
        fmt.Println("Branch Name(After): ", res.branch)

        // Calling the show_2 method
        // (value method) with a pointer
        (&res).show_2()
        fmt.Println("Author's name(After): ", res.name)
    }

    ```

** Output **
    ```console
    Branch Name(Before):  CSE
    Branch Name(After):  ECE
    Author's name(Before) :  Gourav
    Author's name(After):  Sona
    ```

* [Reference](https://www.geeksforgeeks.org/methods-in-golang/)

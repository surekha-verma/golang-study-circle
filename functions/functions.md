# Functions and Methods

## Difference between Functions and Methods
[Reference](https://www.geeksforgeeks.org/methods-in-golang/)

* Go language support methods.
* Go methods are similar to Go function with one difference, i.e, the method contains a receiver argument in it

Difference Between Method and Function
        ** Method **				                            ** Function **
It contains a receiver.								It does not contain a receiver.

Methods of the same name but different
types can be defined in the program.				Functions of the same name but different type are

not allowed to be defined in the program.
It cannot be used as a first-order object.			It can be used as first-order objects and can be passed


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

With the help of the receiver argument, the method can access the properties of the receiver. Here, the receiver can be of struct type or non-struct type.



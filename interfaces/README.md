# Interfaces

An interface is two things: it is a set of methods, but it is also a type.

The primary task of an interface is to provide only function signatures consisting of the name, input arguments and return types. It is up to a type (e.g. struct type) to implement functions.

```console
type <interface_name> interface {
    <method>(<input_arguments>) <return_type>
}
```

Interfaces you use duck typing like you would in a purely dynamic language like Python, but compile time checked.
Duck typing 

Go's interfaces — are static, checked at compile time, dynamic when asked for.


## Go Methods
Go does not have classes. However, you can define methods on types.
A method is a function with a special receiver argument.

The parenthesis before the function name is the Go way of defining the object on which these functions will operate
With the help of the receiver argument, the method can access the properties of the receiver. 

Here, the receiver can be of struct type or non-struct type.


### Interfaces set of Methods and also a Type.

```console
package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width, Height float64
}

type Circle struct {
    Radius float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func getArea(shape Shape) {

    fmt.Println(shape.Area())
}

func main() {

    r := Rectangle{Width: 7, Height: 8}
    c := Circle{Radius: 5}

    getArea(r)
    getArea(c)
}
```

You are allowed to create a method with a pointer receiver.

Go method can accept both value and pointer, whether it is defined with pointer or value receiver.


### Go Interface slice

```console
package main

import (
    "fmt"
)

type Animal interface {
    Sound() string
}

type Dog struct {
}

func (d Dog) Sound() string {
    return "Woof!"
}

type Cat struct {
}

func (c Cat) Sound() string {
    return "Meow!"
}

type Cow struct {
}

func (l Cow) Sound() string {
    return "Moo!"
}

func main() {

    animals := []Animal{Dog{}, Cat{}, Cow{}}

    for _, animal := range animals {

        fmt.Println(animal.Sound())
    }
}
```

### Go Stringer interface
The Stringer interface is defined in the fmt package. Its String function is invoked when a type is passed to any of the print functions. We can customize the output message of our own types.

```console
type Stringer interface {
    String() string
}
```

```console
package main

import (
    "fmt"
)

type User struct {
    Name       string
    Occupation string
}

func (u User) String() string {

    return fmt.Sprintf("%s is a(n) %s", u.Name, u.Occupation)
}

func main() {

    u1 := User{"John Doe", "gardener"}
    u2 := User{"Roger Roe", "driver"}

    fmt.Println(u1)
    fmt.Println(u2)
}
```

### Go empty interface interface{}
Go interface{} is an empty interface; all types in Go satisfy the empty interface. Any type can be assigned to a variable declared with empty interface.

```console
package main

import (
    "fmt"
)

type Body struct {
    Msg interface{}
}

func main() {

    b := Body{"Hello there"}
    fmt.Printf("%#v %T\n", b.Msg, b.Msg)

    b.Msg = 5
    fmt.Printf("%#v %T\n", b.Msg, b.Msg)
}
```

### Array of empty Interfaces

Map with string type keys and interface{} for values

```console
package main

import (
    "fmt"
)

func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"stanley", "david", "oscar"}
    /*vals := make([]interface{}, len(names))
    for i, v := range names {
        vals[i] = v
    }*/
    PrintAll(vals)
}

```

### Go type check with interfaces - runtime type check
Go interfaces and its duck types are checked at compile time, but we can also check for the type of interface at runtime. 

```console
package main

import "fmt"

type User struct {
    Name string
}

func checkType(a interface{}) {

    switch a.(type) {

    case int:
        fmt.Println("Type: int, Value:", a.(int))
    case string:
        fmt.Println("Type: string, Value:", a.(string))
    case float64:
        fmt.Println("Type: float64, Value:", a.(float64))
    case User:
        fmt.Println("Type: User, Value:", a.(User))
    default:
        fmt.Println("unknown type")
    }
}

func main() {

    checkType(4)
    checkType("falcon")
    checkType(User{"John Doe"})
    checkType(7.9)
    checkType(true)
}
```
### Go assertion on interface
The type assertion x.(T) asserts that the concrete value stored in x is of type T, and that x is not nil.

```console
package main

import (
    "fmt"
)

func main() {

    var val interface{} = "falcon"

    r, ok := val.(string)
    fmt.Println(r, ok)

    r2, ok2 := val.(int)
    fmt.Println(r2, ok2)

    r3 := val.(string)
    fmt.Println(r3)

    //r4 := val.(int)
    //fmt.Println(r4)
}
```


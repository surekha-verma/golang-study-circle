# Interfaces

An interface is two things: it is a set of methods, but it is also a type.

The primary task of an interface is to provide only function signatures consisting of the name, input arguments and return types. It is up to a type (e.g. struct type) to implement functions.

Interfaces are also reffered to as exposed APIs or contracts. If a type implements the functions (the APIs) of a sortable interface, we know that it can be sorted; it adheres to the contract of being sortable.

```console
type Shape interface {
    Area() float64
}
```

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

### Go interface{}
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

map with string type keys and interface{} for values

```console

```

### Go type switch
A type switch is used to compare the concrete type of an interface with the multiple types provide in the case statements.

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
### Go type assertion
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



Go's interfaces — are static, checked at compile time, dynamic when asked for.

Go's interfaces let you use duck typing like you would in a purely dynamic language like Python, but compile time checked.
Duck typing 


## Interfaces are set of Methods and also a Type.
What is a method?

Go does not have classes. However, you can define methods on types.
A method is a function with a special receiver argument.

The parenthesis before the function name is the Go way of defining the object on which these functions will operate
With the help of the receiver argument, the method can access the properties of the receiver. 

Here, the receiver can be of struct type or non-struct type.

## Become a type which behaves like an interface
Struct type receiver

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

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
    return "Design patterns!"
}

func main() {
    animals := []Animal{Dog{}, Cat{}, JavaProgrammer{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```

You are allowed to create a method with a pointer receiver.

```console
type Animal interface {
    Speak() string
}

type Dog struct {
}

func (d *Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    animals := []Animal{new(Dog), Cat{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```


Go method can accept both value and pointer, whether it is defined with pointer or value receiver.

```console
type Animal interface {
    Speak() string
}

type Dog struct {
}

func (d *Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    animals := []Animal{new(Dog), new(Cat)}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```

## The interface{} type - empty Interfaces

The interface{} type is the interface that has no methods. Since there is no implements keyword, all types implement at least zero methods, and satisfying an interface is done automatically, all types satisfy the empty interface.

That means that if you write a function that takes an interface{} value as a parameter, you can supply that function with any value.  This function will accept any paramter whatsoever

```console
func DoSomething(val interface{}) {
    fmt.Println(val);
}

func main() {
    names := "stanley"

    DoSomething(names);
}

```

All values have exactly one type at runtime, and v’s one static type is interface{}.

## Checked at compile time

In Go, Interfaces are checked at compile time. If the signature(behaviour) matches to any interface.

## Can be checked at runtime too

```console
type Stringer interface {
    String() string
}

func ToString(any interface{}) string {
    if v, ok := any.(Stringer); ok {
        return v.String()
    }
    switch v := any.(type) {
    case int:
        return strconv.Itoa(v)
    case float32:
        return fmt.Sprintf("%f", v)
    }
    return "???"
}

func main() {
    fmt.Println(ToString("Surekha"));
    fmt.Println(ToString(123));
    fmt.Println(ToString(10.10));
}



```


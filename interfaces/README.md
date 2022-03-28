# Interfaces

An interface is two things: it is a set of methods, but it is also a type. 

```console
type Animal interface {
    Speak() string
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
}


    ```


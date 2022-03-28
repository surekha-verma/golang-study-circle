package interfaces

/**
An interface is two things: it is a set of methods, but it is also a type.
**/
import (
	"fmt"
	"os"
	"strconv"
)

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

func PrintMethods() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

/**
Go's interfaces let you use duck typing like you would in a purely dynamic language like Python
But still have the compiler catch obvious mistakes

The code that calls ReadAndClose can pass a value of any type as long as it has Read and Close methods with the right signatures.
You get error at complie time and not runtime, unlike other programming languages.

**/
type ReadCloser interface {
	Read(b []byte) (n int, err os.SyscallError)
	Close()
}

//The code that calls ReadAndClose can pass a value of any type as long as it has Read and Close methods with the right signatures.
//You get error at complie time and not runtime, unlike other programming languages.
func ReadAndClose(r ReadCloser, buf []byte) (n int) {
	var err os.SyscallError
	var nr int
	for len(buf) > 0 && err.Error() == "" {
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	r.Close()
	return
}

/**
Interfaces aren't restricted to static checking, though.
You can check dynamically whether a particular interface value has an additional method.

**/

type Stringer interface {
	String() string
}

//The value any has static type interface{}, meaning no guarantee of any methods at all: it could contain any type.
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

//let's consider a 64-bit integer type with a String method that prints the value in binary and a trivial Get method
//A value of type Binary can be passed to ToString, which will format it using the String method, even though the program never says that Binary intends to implement Stringer.
type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

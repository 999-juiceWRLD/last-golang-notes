# Some Golang Notes

## If Statements

```pseudo
if INITIAL_STATEMENT; CONDITION {
}
```

This is a syntactic sugar that Go offers. Instead this:

```golang
length := getLength(email)
if length < 1 {
    fmt.Println("email is invalid")
}
```

We can do:

```golang
if length := getLength(email); length < 1 {
    fmt.Println("email is invalid")
}
```

## Functions

After function parentheses, we can define what this function returns.

```golang
func sub(x int, y int) int {
    return x - y
}
```

When multiple arguments are of the same type, then the type can declared only after the last one.

```golang
func add(x, y int) int {
    return x + y
}
```

## Passing Variables by Value

Variables in Go are passed by value (except for some data types). "Pass by value" means that when a variable is passed into a function, that function receives the *copy* of that variable. The function is unable to mutate the caller's data.

## Anonymous Structs

An anonymous struct is a normal struct, but it is defined without a name and therefore cannot be referenced elsewhere in the code.

```golang
myCar := struct {
    Make    string
    Model   string
} {
    Make:   "tesla"
    Model:  "model 3"
}
```

## Embedded Structs

Embedded structs are different than nested structs. Although that Go is not an object-oriented language, embedded structs provide a kind of *data-only* inheritance.

An embedded struct's fields are accessed at the top-level, unlike nested structs.

```golang
type car struct {
    make    string
    model   string
}

type truct struct {
    car
    bedSize int
}

lanesTruck := truck{
    bedSize: 10,
    car: car{
        make: "toyota",
        model: "camry"
    }
}

fmt.Println(lanesTruck.bedSize)
fmt.Println(lanesTruck.make)
fmt.Println(lanesTruck.model)
```

## Struct Methods

```golang
type rect struct {
    width   int
    height  int
}

func (r rect) area() int {
    return r.width * r.heigth
}
```

## Interfaces

Interfaces are implemented *implicitly*. A type never declares that it implements a given interface. If an interface exists and a type has the proper methods defined, then the type automatically fulfills that interface.

```golang
type employee interface {
    getName()   string
    getSalary() int
}

type contractor struct {
    name            string
    hourlyPay       int
    hoursPerYear    int
}

func (c contractor) getName() string {
    return c.name
}

func (c contractor) getSalary() int {
    return c.hourlyPay * c.hoursPerYear
}

cont := contractor{ "jackson", 35, 160 * 12 + 60}
name := cont.getName()
sal := cont.getSalary()
```

## Loops

Basic loop is written in in standard C-like syntax:

```pseudo
for INITIAL; CONDITION; AFTER {
    // do something
}
```

For example:

```golang
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

Loops can omit sections of a for loop. In order to run a loop forever:

```pseudo
for INITIAL; ; AFTER {
    // do something
}

or

for {
    // do something
}
```

There are no while loops in Go. It's declaration is the same as second loop above.

```pseudo
for {
    // ..
}

// or

for CONDITION {
    // do something while CONDITION is true
}
```

## Arrays

To declare arrays in Go:

```golang
var myArr [5]float32

// or

primes := [5]int {2, 3, 5, 7, 11}
```

## Slices

Slices are way more convenient than regular arrays in Go, they are flexible as there are some built-in methods such as `append()`, `copy()` and `delete()`. The declaration of a slice literal is:

```golang
mySlice1 := []int // or
mySlice2 := []int{}
```

Most of the time, we don't need to think about the underlying array of a slice. We can create a new slice using the `make()` function. The `make()` function is a built-in function that is used to allocate and initializes an object of type slice, map, or chan (only).

```golang
// func make([]T, len, capacity) []T
mySlice := make([]int, 5, 10)
mySlice := make([]int, 5) // capacity arg is usually omitted
```

## Variadic Functions

can be called with any number of trailing arguments. For example, `fmt.Println` and `append()` are a common variadic function.

```golang
func sum(nums ...int) {
    num := 0;
    for i := 0; i < len(numbers); i++ {
        num += numbers[i];
    }

    return num;
}
```

## Key Types [Maps]

Any type can be used as the *value* in a map, but *keys* are more restrictive. Map keys may be of any type that is comparable. The comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types. Notably absent from the list are slices, maps, and functions; these types cannot be compared using `==`, and may not be used as map keys.

For example, this map of maps could be used to tally web page hits by country:

```golang
hits := make(map[string]map[string]int)
```

and

```golang
n := hits["/doc/"]["au"]
```

Unfortunately, this approach becomes unwieldy when adding data, as for any given outer key you must check if the inner map exists, and create it if needed:

```golang
func add(m map[string]map[string]int, path, country string) {
    mm, ok := m[path]
    if !ok {
        mm = make(map[string]int)
        m[path] == m
    }
    mm[country]++
}
add(hits, "/doc/", "au")
```

However, a design that uses a single map with a struct key does away with all that complexity:

```golang
type Key struct {
    Path, Country string
}

hits := make(map[Key]int)

hits[Key{"/", "vn"}]++
```

## Closures

A closure is a function that references variables from outside it's own function body.

```golang
func concatter() func(string) string {
    doc := ""
    return func(word string) string {
        doc += word + " "
        return doc;
    }
}

aggregator := concatter();
aggregator("mr.");
aggregator("beast");
```

## Local Development

By convention, a package's name is the same as the last element of its import path. For instance, the `math/rand` package comprises files that begin with:

```golang
package rand
```

A directory of Go code can have **at most** one package. All `.go` files in a single directory must all belong to the same package. If they don't, an error will be thrown by the compiler. This is true for main and library packages alike.

## Modules

Go programs are organized into *packages*. A package is a directory of Go code that's all compiled together. Functions, types, variables and constants defined in one source file are visible to **all other source files withing the same package (directory)**.

A *repository* contains one or more *modules*. A module is a collection of Go packages that are released together.

A Go repository typically contains only one module, located at the root of the repository. A file named `go.mod` at the root of a project declares the module. It contains:

- The module path
- The version of the Go language your project requires
- Optionally, any external package dependencies your project has

Once inside your personal workspace, create a new directory and enter it:

```bash
mkdir hellogo
cd hellogo
```

Inside the directory declare your module's name:

```bash
go mod init {REMOTE}/{USERNAME}/hellogo
```

where `{REMOTE}` is your preferred remote source provider (i.e. `github.com`) and `{USERNAME}` is your Git usename. If you don't use a remote provider yet, just use `example.com/username/hellogo`.

## Go Run

Conventionally, the file in the `main` package that contains the `main()` function is called `main.go`.

## Common Abbreviations in Go

```golang
var s string        // string
var i int           // index
var num int         // number
var msg string      // message
var v string        // value
var val string      // value
var fv string       // flag value
var err error       // error value
var args []string   // arguments
var seen bool       // has (or is) seen?
var parsed bool     // parsing ok?
var buf byte[]      // buffer
var off int         // offset
var op int          // operation
var opRead int      // read operation
var l int           // length
var n int           // number of
var m int           // another number
var c int           // capacity
var c int           // character
var a int           // array
var r rune          // rune
var sep string      // separator
var source int      // source
var dst int         // destionation
var b byte          // byte
var b []byte        // buffer
var buf []byte      // buffer
var w io.Writer     // writer
var r io.Reader     // reader
var pos int         // position

// list goes on...
```

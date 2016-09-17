package main

import "fmt"
//import "math"
//import "errors"
import "time"

// const s string = "constant"


// // Here's a function that takes two `int`s and returns
// // their sum as an `int`.
// func plus(a int, b int) int {

//     // Go requires explicit returns, i.e. it won't
//     // automatically return the value of the last
//     // expression.
//     return a + b
// }

// // When you have multiple consecutive parameters of
// // the same type, you may omit the type name for the
// // like-typed parameters up to the final parameter that
// // declares the type.
// func plusPlus(a, b, c int) int {
//     return a + b + c
// }

// func vals() (int, int) {
//     return 3, 7
// }


// // Here's a function that will take an arbitrary number
// // of `ints` as arguments.
// func sum(nums ...int) {
//     fmt.Print(nums, " ")
//     total := 0
//     for _, num := range nums {
//         total += num
//     }
//     fmt.Println(total)
// }

// func intSeq() func() int {
//     i := 0
//     return func() int {
//         i += 1
//         return i
//     }
// }

// func fact(n int) int {
//     if n == 0 {
//         return 1
//     }
//     return n * fact(n-1)
// }

// // We'll show how pointers work in contrast to values with
// // 2 functions: `zeroval` and `zeroptr`. `zeroval` has an
// // `int` parameter, so arguments will be passed to it by
// // value. `zeroval` will get a copy of `ival` distinct
// // from the one in the calling function.
// func zeroval(ival int) {
//     ival = 0
// }

// // `zeroptr` in contrast has an `*int` parameter, meaning
// // that it takes an `int` pointer. The `*iptr` code in the
// // function body then _dereferences_ the pointer from its
// // memory address to the current value at that address.
// // Assigning a value to a dereferenced pointer changes the
// // value at the referenced address.
// func zeroptr(iptr *int) {
//     *iptr = 0
// }

	
// type person struct {
//     name string
//     age  int
// }


// type rect struct {
//     width, height int
// }

// // This `area` method has a _receiver type_ of `*rect`.
// func (r *rect) area() int {
//     return r.width * r.height
// }

// // Methods can be defined for either pointer or value
// // receiver types. Here's an example of a value receiver.
// func (r rect) perim() int {
//     return 2*r.width + 2*r.height
// }


// // Here's a basic interface for geometric shapes.
// type geometry interface {
//     area() float64
//     perim() float64
// }

// // For our example we'll implement this interface on
// // `rect` and `circle` types.
// type rect struct {
//     width, height float64
// }
// type circle struct {
//     radius float64
// }

// // To implement an interface in Go, we just need to
// // implement all the methods in the interface. Here we
// // implement `geometry` on `rect`s.
// func (r rect) area() float64 {
//     return r.width * r.height
// }
// func (r rect) perim() float64 {
//     return 2*r.width + 2*r.height
// }

// // The implementation for `circle`s.
// func (c circle) area() float64 {
//     return math.Pi * c.radius * c.radius
// }
// func (c circle) perim() float64 {
//     return 2 * math.Pi * c.radius
// }

// // If a variable has an interface type, then we can call
// // methods that are in the named interface. Here's a
// // generic `measure` function taking advantage of this
// // to work on any `geometry`.
// func measure(g geometry) {
//     fmt.Println(g)
//     fmt.Println(g.area())
//     fmt.Println(g.perim())
// }



// // By convention, errors are the last return value and
// // have type `error`, a built-in interface.
// func f1(arg int) (int, error) {
//     if arg == 42 {

//         // `errors.New` constructs a basic `error` value
//         // with the given error message.
//         return -1, errors.New("can't work with 42")

//     }

//     // A nil value in the error position indicates that
//     // there was no error.
//     return arg + 3, nil
// }

// // It's possible to use custom types as `error`s by
// // implementing the `Error()` method on them. Here's a
// // variant on the example above that uses a custom type
// // to explicitly represent an argument error.
// type argError struct {
//     arg  int
//     prob string
// }

// func (e *argError) Error() string {
//     return fmt.Sprintf("%d - %s", e.arg, e.prob)
// }

// func f2(arg int) (int, error) {
//     if arg == 42 {

//         // In this case we use `&argError` syntax to build
//         // a new struct, supplying values for the two
//         // fields `arg` and `prob`.
//         return -1, &argError{arg, "can't work with it"}
//     }
//     return arg + 3, nil
// }

// func f(from string) {
//     for i := 0; i < 3; i++ {
//         fmt.Println(from, ":", i)
//     }
// }

// // This is the function we'll run in a goroutine. The
// // `done` channel will be used to notify another
// // goroutine that this function's work is done.
// func worker(done chan bool) {
//     fmt.Print("working...")
//     time.Sleep(time.Second)
//     fmt.Println("done")

//     // Send a value to notify that we're done.
//     done <- true
// }

// This `ping` function only accepts a channel for sending
// values. It would be a compile-time error to try to
// receive on this channel.
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// The `pong` function accepts one channel for receives
// (`pings`) and a second for sends (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    // fmt.Printf("This should now work\n")

    // // Strings, which can be added together with `+`.
    // fmt.Println("go" + "lang")

    // // Integers and floats.
    // fmt.Println("1+1 =", 1+1)
    // fmt.Println("7.0/3.0 =", 7.0/3.0)

    // // Booleans, with boolean operators as you'd expect.
    // fmt.Println(true && false)
    // fmt.Println(true || false)
    // fmt.Println(!true)

    // // `var` declares 1 or more variables.
    // var a string = "initial"
    // fmt.Println(a)

    // // You can declare multiple variables at once.
    // var b, c int = 1, 2
    // fmt.Println(b, c)


    // // Variables declared without a corresponding
    // // initialization are _zero-valued_. For example, the
    // // zero value for an `int` is `0`.
    // var e int
    // fmt.Println(e)

    // // The `:=` syntax is shorthand for declaring and
    // // initializing a variable, e.g. for
    // // `var f string = "short"` in this case.
    // f := "short"
    // fmt.Println(f)


    // fmt.Println(s)

    // // A `const` statement can appear anywhere a `var`
    // // statement can.
    // const n = 500000000

    // // Constant expressions perform arithmetic with
    // // arbitrary precision.
    // const d = 3e20 / n
    // fmt.Println(d)

    // // A numeric constant has no type until it's given
    // // one, such as by an explicit cast.
    // fmt.Println(int64(d))

    // // A number can be given a type by using it in a
    // // context that requires one, such as a variable
    // // assignment or function call. For example, here
    // // `math.Sin` expects a `float64`.
    // fmt.Println(math.Sin(n))

    // fmt.Printf("******* for loop *******\n")

    // // The most basic type, with a single condition.
    // i := 1
    // for i <= 3 {
    //     fmt.Println(i)
    //     i = i + 1
    // }

    // // A classic initial/condition/after `for` loop.
    // for j := 7; j <= 9; j++ {
    //     fmt.Println(j)
    // }

    // // `for` without a condition will loop repeatedly
    // // until you `break` out of the loop or `return` from
    // // the enclosing function.
    // for {
    //     fmt.Println("loop")
    //     break
    // }


    // fmt.Printf("******* if statement *******\n")

    // // Here's a basic example.
    // if 7%2 == 0 {
    //     fmt.Println("7 is even")
    // } else {
    //     fmt.Println("7 is odd")
    // }

    // // You can have an `if` statement without an else.
    // if 8%4 == 0 {
    //     fmt.Println("8 is divisible by 4")
    // }

    // // A statement can precede conditionals; any variables
    // // declared in this statement are available in all
    // // branches.
    // if num := 9; num < 0 {
    //     fmt.Println(num, "is negative")
    // } else if num < 10 {
    //     fmt.Println(num, "has 1 digit")
    // } else {
    //     fmt.Println(num, "has multiple digits")
    // }


    // fmt.Printf("******* switch statement *******\n")


    // // You can use commas to separate multiple expressions
    // // in the same `case` statement. We use the optional
    // // `default` case in this example as well.
    // switch time.Now().Weekday() {
    // case time.Saturday, time.Sunday:
    //     fmt.Println("it's the weekend")
    // default:
    //     fmt.Println("it's a weekday")
    // }

    // // `switch` without an expression is an alternate way
    // // to express if/else logic. Here we also show how the
    // // `case` expressions can be non-constants.
    // t := time.Now()
    // switch {
    // case t.Hour() < 12:
    //     fmt.Println("it's before noon")
    // default:
    //     fmt.Println("it's after noon")
    // }

    // fmt.Printf("******* array *******\n")

    // // Here we create an array `a` that will hold exactly
    // // 5 `int`s. The type of elements and length are both
    // // part of the array's type. By default an array is
    // // zero-valued, which for `int`s means `0`s.
    // var arr [5]int
    // fmt.Println("emp:", arr)

    // // We can set a value at an index using the
    // // `array[index] = value` syntax, and get a value with
    // // `array[index]`.
    // arr[4] = 100
    // fmt.Println("set:", arr)
    // fmt.Println("get:", arr[4])

    // // The builtin `len` returns the length of an array.
    // fmt.Println("len:", len(arr))

    // // Use this syntax to declare and initialize an array
    // // in one line.
    // bArr := [5]int{1, 2, 3, 4, 5}
    // fmt.Println("dcl:", bArr)

    // // Array types are one-dimensional, but you can
    // // compose types to build multi-dimensional data
    // // structures.
    // var twoD [2][3]int
    // for i := 0; i < 2; i++ {
    //     for j := 0; j < 3; j++ {
    //         twoD[i][j] = i + j
    //     }
    // }
    // fmt.Println("2d: ", twoD)

    // fmt.Printf("******* Slices *******\n")


    // // Unlike arrays, slices are typed only by the
    // // elements they contain (not the number of elements).
    // // To create an empty slice with non-zero length, use
    // // the builtin `make`. Here we make a slice of
    // // `string`s of length `3` (initially zero-valued).
    // s := make([]string, 3)
    // fmt.Println("emp:", s)

    // // We can set and get just like with arrays.
    // s[0] = "a"
    // s[1] = "b"
    // s[2] = "c"
    // fmt.Println("set:", s)
    // fmt.Println("get:", s[2])

    // // `len` returns the length of the slice as expected.
    // fmt.Println("len:", len(s))

    // // In addition to these basic operations, slices
    // // support several more that make them richer than
    // // arrays. One is the builtin `append`, which
    // // returns a slice containing one or more new values.
    // // Note that we need to accept a return value from
    // // append as we may get a new slice value.
    // s = append(s, "d")
    // s = append(s, "e", "f")
    // fmt.Println("apd:", s)

    // // Slices can also be `copy`'d. Here we create an
    // // empty slice `c` of the same length as `s` and copy
    // // into `c` from `s`.
    // c := make([]string, len(s))
    // copy(c, s)
    // fmt.Println("cpy:", c)

    // // Slices support a "slice" operator with the syntax
    // // `slice[low:high]`. For example, this gets a slice
    // // of the elements `s[2]`, `s[3]`, and `s[4]`.
    // l := s[2:5]
    // fmt.Println("sl1:", l)

    // // This slices up to (but excluding) `s[5]`.
    // l = s[:5]
    // fmt.Println("sl2:", l)

    // // And this slices up from (and including) `s[2]`.
    // l = s[2:]
    // fmt.Println("sl3:", l)

    // // We can declare and initialize a variable for slice
    // // in a single line as well.
    // t := []string{"g", "h", "i"}
    // fmt.Println("dcl:", t)

    // // Slices can be composed into multi-dimensional data
    // // structures. The length of the inner slices can
    // // vary, unlike with multi-dimensional arrays.
    // twoD := make([][]int, 3)
    // for i := 0; i < 3; i++ {
    //     innerLen := i + 1
    //     twoD[i] = make([]int, innerLen)
    //     for j := 0; j < innerLen; j++ {
    //         twoD[i][j] = i + j
    //     }
    // }
    // fmt.Println("2d: ", twoD)


    // fmt.Printf("******* Map *******\n")
    // // map is an object 

    // // To create an empty map, use the builtin `make`:
    // // `make(map[key-type]val-type)`.
    // m := make(map[string]int)

    // // Set key/value pairs using typical `name[key] = val`
    // // syntax.
    // m["k1"] = 7
    // m["k2"] = 13

    // // Printing a map with e.g. `Println` will show all of
    // // its key/value pairs.
    // fmt.Println("map:", m)

    // // Get a value for a key with `name[key]`.
    // v1 := m["k1"]
    // fmt.Println("v1: ", v1)

    // // The builtin `len` returns the number of key/value
    // // pairs when called on a map.
    // fmt.Println("len:", len(m))

    // // The builtin `delete` removes key/value pairs from
    // // a map.
    // delete(m, "k2")
    // fmt.Println("map:", m)

    // // The optional second return value when getting a
    // // value from a map indicates if the key was present
    // // in the map. This can be used to disambiguate
    // // between missing keys and keys with zero values
    // // like `0` or `""`. Here we didn't need the value
    // // itself, so we ignored it with the _blank identifier_
    // // `_`.
    // _, prs := m["k2"]
    // fmt.Println("prs:", prs)

    // // You can also declare and initialize a new map in
    // // the same line with this syntax.
    // n := map[string]int{"foo": 1, "bar": 2}
    // fmt.Println("map:", n)

	// fmt.Printf("******* Range *******\n")

	// // Here we use `range` to sum the numbers in a slice.
	// // Arrays work like this too.
	// nums := []int{2, 3, 4}
	// sum := 0
	// for _, num := range nums {
	//     sum += num
	// }
	// fmt.Println("sum:", sum)

	// // `range` on arrays and slices provides both the
	// // index and value for each entry. Above we didn't
	// // need the index, so we ignored it with the
	// // blank identifier `_`. Sometimes we actually want
	// // the indexes though.
	// for i, num := range nums {
	//     if num == 3 {
	//         fmt.Println("index:", i)
	//     }
	// }

	// // `range` on map iterates over key/value pairs.
	// kvs := map[string]string{"a": "apple", "b": "banana"}
	// for k, v := range kvs {
	//     fmt.Printf("%s -> %s\n", k, v)
	// }

	// // `range` on strings iterates over Unicode code
	// // points. The first value is the starting byte index
	// // of the `rune` and the second the `rune` itself.
	// for i, c := range "go" {
	//     fmt.Println(i, c)
	// }


	// fmt.Printf("******* Function *******\n")

	// // Call a function just as you'd expect, with
	// // `name(args)`.
	// res := plus(1, 2)
	// fmt.Println("1+2 =", res)

	// res = plusPlus(1, 2, 3)
	// fmt.Println("1+2+3 =", res)


	// fmt.Printf("******* multiple variable function *******\n")

	// // Here we use the 2 different return values from the
	// // call with _multiple assignment_.
	// a, b := vals()
	// fmt.Println(a)
	// fmt.Println(b)

	// // If you only want a subset of the returned values,
	// // use the blank identifier `_`.
	// _, c := vals()
	// fmt.Println(c)

	// fmt.Printf("******* Variadic functions *******\n")

	// // Variadic functions can be called in the usual way
	// // with individual arguments.
	// sum(1, 2)
	// sum(1, 2, 3)

	// // If you already have multiple args in a slice,
	// // apply them to a variadic function using
	// // `func(slice...)` like this.
	// nums := []int{1, 2, 3, 4}
	// sum(nums...)

	// fmt.Printf("******* Closure *******\n")

	// // We call `intSeq`, assigning the result (a function)
	// // to `nextInt`. This function value captures its
	// // own `i` value, which will be updated each time
	// // we call `nextInt`.
	// nextInt := intSeq()

	// // See the effect of the closure by calling `nextInt`
	// // a few times.
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())

	// // To confirm that the state is unique to that
	// // particular function, create and test a new one.
	// newInts := intSeq()
	// fmt.Println(newInts())


	// fmt.Printf("******* Recursion *******\n")

	// fmt.Println(fact(7))

	// // fmt.Printf("******* Pointer *******\n")
	// i := 1
	// fmt.Println("initial:", i)

	// zeroval(i)
	// fmt.Println("zeroval:", i)

	// // The `&i` syntax gives the memory address of `i`,
	// // i.e. a pointer to `i`.
	// zeroptr(&i)
	// fmt.Println("zeroptr:", i)

	// // Pointers can be printed too.
	// fmt.Println("pointer:", &i)

	// fmt.Printf("******* Struct *******\n")
	// // This syntax creates a new struct.
	// fmt.Println(person{"Bob", 20})

	// // You can name the fields when initializing a struct.
	// fmt.Println(person{name: "Alice", age: 30})

	// // Omitted fields will be zero-valued.
	// fmt.Println(person{name: "Fred"})

	// // An `&` prefix yields a pointer to the struct.
	// fmt.Println(&person{name: "Ann", age: 40})

	// // Access struct fields with a dot.
	// s := person{name: "Sean", age: 50}
	// fmt.Println(s.name)

	// // You can also use dots with struct pointers - the
	// // pointers are automatically dereferenced.
	// sp := &s
	// fmt.Println(sp.age)

	// // Structs are mutable.
	// sp.age = 51
	// fmt.Println(sp.age)

	// // fmt.Printf("******* Method *******\n")

	// r := rect{width: 10, height: 5}

	//   // Here we call the 2 methods defined for our struct.
	//   fmt.Println("area: ", r.area())
	//   fmt.Println("perim:", r.perim())

	//   // Go automatically handles conversion between values
	//   // and pointers for method calls. You may want to use
	//   // a pointer receiver type to avoid copying on method
	//   // calls or to allow the method to mutate the
	//   // receiving struct.
	//   rp := &r
	//   fmt.Println("area: ", rp.area())
	//   fmt.Println("perim:", rp.perim())


	// // // fmt.Printf("******* Interface *******\n")
	// r := rect{width: 3, height: 4}
	// c := circle{radius: 5}

	// // The `circle` and `rect` struct types both
	// // implement the `geometry` interface so we can use
	// // instances of
	// // these structs as arguments to `measure`.
	// measure(r)
	// measure(c)


	//  fmt.Printf("******* Error *******\n")

	// // The two loops below test out each of our
	// // error-returning functions. Note that the use of an
	// // inline error check on the `if` line is a common
	// // idiom in Go code.
	// for _, i := range []int{7, 42} {
	//     if r, e := f1(i); e != nil {
	//         fmt.Println("f1 failed:", e)
	//     } else {
	//         fmt.Println("f1 worked:", r)
	//     }
	// }
	// for _, i := range []int{7, 42} {
	//     if r, e := f2(i); e != nil {
	//         fmt.Println("f2 failed:", e)
	//     } else {
	//         fmt.Println("f2 worked:", r)
	//     }
	// }

	// // If you want to programmatically use the data in
	// // a custom error, you'll need to get the error  as an
	// // instance of the custom error type via type
	// // assertion.
	// _, e := f2(42)
	// if ae, ok := e.(*argError); ok {
	//     fmt.Println(ae.arg)
	//     fmt.Println(ae.prob)
	// }


	// fmt.Printf("******* Go Routines *******\n")

	//  // Suppose we have a function call `f(s)`. Here's how
 //    // we'd call that in the usual way, running it
 //    // synchronously.
 //    f("direct")

 //    // To invoke this function in a goroutine, use
 //    // `go f(s)`. This new goroutine will execute
 //    // concurrently with the calling one.
 //    go f("goroutine")

 //    // You can also start a goroutine for an anonymous
 //    // function call.
 //    go func(msg string) {
 //        fmt.Println(msg)
 //    }("going")

 //    // Our two function calls are running asynchronously in
 //    // separate goroutines now, so execution falls through
 //    // to here. This `Scanln` code requires we press a key
 //    // before the program exits.
 //    var input string
 //    fmt.Scanln(&input)
 //    fmt.Println("done")

	// fmt.Printf("******* Channel *******\n")

	// // Create a new channel with `make(chan val-type)`.
	// // Channels are typed by the values they convey.
	// messages := make(chan string)

	// // _Send_ a value into a channel using the `channel <-`
	// // syntax. Here we send `"ping"`  to the `messages`
	// // channel we made above, from a new goroutine.
	// go func() { messages <- "ping" }()

	// // The `<-channel` syntax _receives_ a value from the
	// // channel. Here we'll receive the `"ping"` message
	// // we sent above and print it out.
	// msg := <-messages
	// fmt.Println(msg)

	// fmt.Printf("******* Channel *******\n")

	// // Here we `make` a channel of strings buffering up to
	// // 2 values.
	// messages := make(chan string, 2)

	// // Because this channel is buffered, we can send these
	// // values into the channel without a corresponding
	// // concurrent receive.
	// messages <- "buffered"
	// messages <- "channel"

	// // Later we can receive these two values as usual.
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)

	// fmt.Printf("******* synchronize execution across goroutines *******\n")

	// // Start a worker goroutine, giving it the channel to
	// // notify on.
	// done := make(chan bool, 1)
	// go worker(done)

	// // Block until we receive a notification from the
	// // worker on the channel.
	// <-done

	// pings := make(chan string, 1)
 //  pongs := make(chan string, 1)
 //  ping(pings, "passed message")
 //  pong(pings, pongs)
 //  fmt.Println(<-pongs)

  // fmt.Printf("******* Select *******\n")

  // // For our example we'll select across two channels.
  // c1 := make(chan string)
  // c2 := make(chan string)

  // // Each channel will receive a value after some amount
  // // of time, to simulate e.g. blocking RPC operations
  // // executing in concurrent goroutines.
  // go func() {
  //     time.Sleep(time.Second * 1)
  //     c1 <- "one"
  // }()
  // go func() {
  //     time.Sleep(time.Second * 2)
  //     c2 <- "two"
  // }()

  // // We'll use `select` to await both of these values
  // // simultaneously, printing each one as it arrives.
  // for i := 0; i < 2; i++ {
  //     select {
  //     case msg1 := <-c1:
  //         fmt.Println("received", msg1)
  //     case msg2 := <-c2:
  //         fmt.Println("received", msg2)
  //     }
  // }

  fmt.Printf("******* Timeout *******\n")

  // For our example, suppose we're executing an external
  // call that returns its result on a channel `c1`
  // after 2s.
  c1 := make(chan string, 1)
  go func() {
      time.Sleep(time.Second * 2)
      c1 <- "result 1"
  }()

  // Here's the `select` implementing a timeout.
  // `res := <-c1` awaits the result and `<-Time.After`
  // awaits a value to be sent after the timeout of
  // 1s. Since `select` proceeds with the first
  // receive that's ready, we'll take the timeout case
  // if the operation takes more than the allowed 1s.
  select {
  case res := <-c1:
      fmt.Println(res)
  case <-time.After(time.Second * 1):
      fmt.Println("timeout 1")
  }

  // If we allow a longer timeout of 3s, then the receive
  // from `c2` will succeed and we'll print the result.
  c2 := make(chan string, 1)
  go func() {
      time.Sleep(time.Second * 2)
      c2 <- "result 2"
  }()
  select {
  case res := <-c2:
      fmt.Println(res)
  case <-time.After(time.Second * 3):
      fmt.Println("timeout 2")
  }

}

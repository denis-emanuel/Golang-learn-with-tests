# Learning GO

Learning GO with testing from <a href="https://quii.gitbook.io/learn-go-with-tests/">this</a> source

- [Learning GO](#learning-go)
  - [**Useful things to remember**](#useful-things-to-remember)
  - [**Folders**](#folders)
    - [**Hello World**](#hello-world)
    - [**Struct, methods & interfaces**](#struct-methods--interfaces)
      - [**Structs**](#structs)
      - [**Methods**](#methods)
        - [Example](#example)
      - [**Interfaces**](#interfaces)
    - [**Pointers & Errors**](#pointers--errors)
    - [**Dependency Injection**](#dependency-injection)
    - [**Concurrency**](#concurrency)
      - [Anonymous function example to start a gorouting](#anonymous-function-example-to-start-a-gorouting)
      - [Anonymous function features](#anonymous-function-features)
      - [Quick note on parallelism and a small problem](#quick-note-on-parallelism-and-a-small-problem)
      - [Another problem on **_DATA RACE_**](#another-problem-on-data-race)
      - [Solution: Channels](#solution-channels)
  - [Select](#select)
    - [Select](#select-1)
      - [defer](#defer)
      - [select](#select-2)
      - [httptest](#httptest)
  - [Sync](#sync)
    - [WaitGroup](#waitgroup)
    - [MUTEX](#mutex)

## **Useful things to remember**

- In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.

<br>

## **Folders**

Folders are listed here in the order they were approached

<br>

### **Hello World**

A first example with testing of a basic function to get familiarized with the Go syntax and the way of writing tests.

The idea is to write tests and assertions first and then refactor the code to meet the test requirements. This pattern will be used for all the code written in Go in this repo

<br>

### **Struct, methods & interfaces**

#### **Structs**

A **struct** is just a named collection of fields where you can store data

```go
type Rectangle struct {
	Width  float64
	Height float64
}
```

#### **Methods**

The syntax for declaring methods is almost the same as functions and that's because they're so similar.

The only difference is the syntax of the method receiver func `(receiverName ReceiverType) MethodName(args)`.

When your method is called on a variable of that type, you get your reference to its data via the `receiverName` variable. In many other programming languages this is done implicitly and you access the receiver via `this`.

##### Example

```go
func (r Rectangle) Area() float64 {
	return 0
}
```

It is a convention in Go to have the receiver variable be the first letter of the type.

```go
r Rectangle
```

<br>

#### **Interfaces**

They allow you to make functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety

For our example we create a `Shape` interface, so if anything that's not a shape is passed in it will not compile.

How does something become a _Shape_? We just tell Go what a _Shape_ is using an interface declaration

```go
type Shape interface {
	Area() float64
}
```

**!!!** Normally in other languages you have to specifically say `MyType Rectangle implements Shape` but in Go's case:

- `Rectangle` has a method called `Area()` that returns a `float64` so it satisfies the `Shape` interface
- Same goes for `Circle`

In Go **_Interface resolution is implicit_**. If the type you pass in matches what the interface is asking for, it will compile.

<br>

### **Pointers & Errors**

`nil` is synonymous with `null` from other programming languages.

Errors can be `nil` if the return type of a function is `error` (which is an interface)

If you see a function that takes arguments or returns values that are interfaces, they can be _nillable_

Just like `null` if you try to access a value that is `nil` it will throw a runtime panic.

**!!!** Make sure to check for nils.

<br>

Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.

<br>

### **Dependency Injection**

- You don't need a framework
- It does not overcomplicate your design
- It facilitates testing
- It allows you to write great, general-purpose functions.

Assume we have this function

```go
func Greet(name string) {
	fmt.Printf("Hello, %s", name)
}
```

How can we test it because `fmt.Printf` prints to stdout and it's pretty hard for us to capture using the testing framework.

We need to be able to **inject** (pass in) the dependency of printing

Our function doesn't need to care where or how the printing happens, so we should accept an interface rather than a concrete type and by doing that we can change the implementation to print something we control so that we can test it.

<br>

### **Concurrency**

To tell go to start a `goroutine` we turn a function call into a `go` statement by putting the `go` in front of it like `go doSomething()`

And because of that we often use _anonymous functions_ when we want to start a goroutine.

<br>

#### Anonymous function example to start a gorouting

```go
results := make(map[string]bool)

for _, url := range urls {
	go func() { //Anonymous function
		results[url] = wc(url)
	}()
}
```

#### Anonymous function features

- they can be executed at the same time they are declared (notice the "()" at the end of the declaration)
- they maintain access to the lexical scope they are defined in

<br>

#### Quick note on parallelism and a small problem

If we run a routine like this in a for loop:

```go
results := make(map[string]bool)

for _, url := range urls {
	go func() {
  	results[url] = wc(url)
	}()
}

return results
```

none of the goroutines in the for loop have enough time to add their result to the `results` map because the `wc` function is too fast for them and it returns an empty array.

So we can add a wait to wait until the goroutines do their work and then return. So we can add this code in between the `anonymous function` call and the `return` statement.

```go
time.Sleep(2 * time.Second)
```

But this option has another flaw. It displays only one result in the `result` map.
The problem is that the `url` is reused for each iteration of the `for` loop.

Each goroutines has a reference to the `url` variable, they don't have their own independent copy. And so, they are all writing the value that `url` has at the end of the iteration.

_How to fix_:

```go
results := make(map[string]bool)

for _, url := range urls {
	go func(u string) {
		results[u] = wc(u)
	}(url)
}

time.Sleep(2 * time.Second)

return results
```

In this solution we give each anonymous function a parameter for the url called `u` and we instantly call them with the `url` as an argument.

Thus we make sure that the value of `u` is fixed as the value of `url` for the iteration of the loop that we're launching the goroutine in.

`u` is a copy value of `url` and so can't be changed.

<br>

#### Another problem on **_DATA RACE_**

Sometimes, when we run our tests, two of the goroutines write to the results map at exactly the same time. Maps in Go don't like this so -> **fatal error**

<br>

#### Solution: Channels

We can coordinate our goroutines using _channels_.  
_Channels_ are a Go data structure that can both receive and send values. These operations allow for communication between different processes.

For example we can think about the communication between the parent process and each of the goroutines.

<br>

## Select

### Select

#### defer

By prefixing a function with _defer_ it will now call that function at the end of the containing function.

For example you may want to close an open port and have that code written whever you opened the port for better readability.

<br>

#### select

It's a construct which helps us synchronize processes really easily and clearly.  
_Select_ lets you wait on multiple channels. The first one to send a value wins and the code underneath the `case` is executed

<br>

#### httptest

A convenient way of creating test servers so you can have realiable and controllable tests

<br>

## Sync

How to make a counter that is safe to use in a concurrent environment.

<br>

### WaitGroup

A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls `Add` to set the number of goroutines to wait for. Then each of the goroutines runs and calls `Done` when finished. At the same time, `Wait` can be used to block until all goroutines have finished.

<br>

### MUTEX

Allows us to add locks to our data.

**!!!** A mutex must not be copied after first use.

A _Mutex_ is a mutual exclusion lock. The zero value for a Murex is an unlocked mutex.

For a good example on how to use this go to [this code](Sync/sync.go) and check the comments for the mutex.

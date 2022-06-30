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

## **Useful things to remember**

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

**!** Normally in other languages you have to specifically say `MyType Rectangle implements Shape` but in Go's case:

- `Rectangle` has a method called `Area()` that returns a `float64` so it satisfies the `Shape` interface
- Same goes for `Circle`

In Go **_Interface resolution is implicit_**. If the type you pass in matches what the interface is asking for, it will compile.

<br>

### **Pointers & Errors**

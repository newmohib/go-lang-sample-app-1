package main

import (
	"fmt"

	mathlib "example.com/math"
)

var (
	a = 20
	b = 30
)

func main() {
	fmt.Println("Hello, World!")

	// mathlib.Sum(1, 2) // argument

	// anonymous function
	// immediatly invoded fuction expression

	func() { // fuction expression
		fmt.Println(a + b)
	}()

	processOperation(2, 3, mathlib.Sum)

	// anonymous function
	sum := call() // fuction expression
	fmt.Println("Print call function value", sum(a, b))

	user()

}

// Higher Order Function
func processOperation(a int, b int, operation func(int, int)) {
	operation(a, b)

}

func call() func(x int, y int) int {
	return mathlib.Add
}

func init() {
	fmt.Println("init function")
}

// struct
type Person struct {
	name string
	age  int
}

func user() {

	person := Person{
		name: "John",
		age:  30,
	}
	printUserDetails(person)

	var person2 Person
	person2.name = "Test User"
	person2.age = 20
	printUserDetails(person2)
	person2.printDetails()

	// Arrays
	arrF()

	// pointer
	pointer()

}

// Recisiver Function

func printUserDetails(user Person) {
	fmt.Println("Name: ", user.name)
	fmt.Println("Age: ", user.age)
}

// Receiver function
// without custome type its not posible

func (user Person) printDetails() {
	fmt.Println("Name: ", user.name)
	fmt.Println("Age: ", user.age)
}

// Aeeys

func arrF() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
}

// pointer or address of memory (ram  )

func pointer() {
	x := 10
	fmt.Println(x)
	fmt.Println(&x) // ampersand & => address of memory

	p := &x // appersand & => address of memory as haxa value

	fmt.Println(" address of memory", p)     // => address of memory as haxa value
	fmt.Println(" value at the address", *p) // value at the address, dereference * => value

	*p = 20
	fmt.Println(" value at the address of x", x)

	// for using pointer in array
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// send pointer address to function as argument
	print(&arr)

	userInfo()

	slice()

	// variadic function can provide any number of arguments
	variadicSum(1,2,3,4,5)

	// Vogus Data types
	vogusDataType()

}

// recive a pointer as parameter
func print(numbers *[5]int) {
	fmt.Println(" receive data by pointer", numbers)
}

// pass by value
// pass by reference

type User struct {
	Name   string
	Age    int
	Salary float64
}

func userInfo() {
	obj := User{ // instance or object
		Name:   "John",
		Age:    30,
		Salary: 50000,
	}

	fmt.Println(obj)

	p := &obj
	fmt.Println(p)
	fmt.Println(*p)

} // pointer, struct, slice, reciver function

// slice
// slice have 3 type or meta data =>> 1: pointer 2: length 3: capacity

// 1: slice from an existing array
// 2: slcie from a slice
// 3: slice literal
// 4: make function with len
// 5: make function with len and capacity
// 6: empty slice or nil slice

func slice() {
	arr := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(arr)

	s := arr[2:5]  // slice
	fmt.Println(s) // c-----e

	t := s[2:4]    // slice
	fmt.Println(t) // e-----f

	u := []int{1, 2, 3, 4, 5} // slice literal
	fmt.Println("slice: ", u, "Len:", len(u), "Cap:", cap(u))

	// slice with make without capacity
	s1 := make([]int, 5) // slice with make
	fmt.Println("slice: ", s1, "Len:", len(s1), "Cap:", cap(s1))

	// slice with make with capacity
	s2 := make([]int, 5, 10) // slice with make
	fmt.Println("slice: ", s2, "Len:", len(s2), "Cap:", cap(s2))

	// empty slice or nil slice
	var s3 []int // empty slice
	s3 = append(s3, 1, 2, 3, 4, 5)

	fmt.Println("slice: ", s3, "Len:", len(s3), "Cap:", cap(s3))

}

// variadic function
// variadic function can provide any number of arguments
func variadicSum(a ...int) {
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}


// Vogus Data Types


func vogusDataType()  {
	var a int8 = -128
	var b int8 = 127
	var x uint8 = 255 // unsigned (0 and only positive numbers)

	var j float32 = 10.3333

	var flag bool = true // 8 bit
	// run
	r := '😄' // alias for int32 (unicode point) -> 32 bits -> 4bytes -> %c

	fmt.Printf("%c\n", r) // character => printf()
	fmt.Printf("%d\n", a) // decimal => printf()
	fmt.Printf("%f\n", j) // floating => printf()
	fmt.Printf("%.2f\n", j) // floating and after decimal will be 2 character => printf()
	fmt.Printf("%v\n", flag) // boolean => printf()
	fmt.Printf("%T\n", b) // print type => printf()
}



func deferFunc()  {
	//Defer
	i := 0

	fmt.Println("first", i) // i = 0

	defer fmt.Println("second", i) // i = 0

	i = i + 1 // or i++

	fmt.Println("third", i) //  i = 1

	return
}
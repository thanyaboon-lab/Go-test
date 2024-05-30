package bambino

import (
	"errors"
	"fmt"
)

func SayHello() {
	num1 := 5
	num2 := 10
	if sumNum := num1 + num2; sumNum >= 10 {
		fmt.Printf("%d \n", sumNum)
	}

	// Array
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	for i := 0; i < len(myArray); i++ {
		fmt.Printf("%d \n", i)
	}

	// Slice
	mySlice := []int{10, 20, 30, 40, 50}
	fmt.Println(mySlice)

	subSlice := mySlice[1:3]
	fmt.Println(subSlice)
	fmt.Println(len(subSlice))
	fmt.Println(cap(subSlice))

	// Convert Array to Slice
	mySlice2 := myArray[:]
	fmt.Println(mySlice2)

	fmt.Println("Say Hello Bambino")
	generateTest()

	// Map
	myMap := make(map[string]int)
	myMap["banana"] = 10
	myMap["orange"] = 8
	myMap["apple"] = 5
	fmt.Println(myMap)

	//Iterate over the map
	for key, value := range myMap {
		fmt.Printf("%s -> %d\n", key, value)
	}

	// Checking if a key exists
	val, ok := myMap["peer"]
	fmt.Println("val:", val)
	fmt.Println("ok:", ok)

	if ok {
		fmt.Println("Peer's value", val)
	} else {
		fmt.Println("Peer not found")
	}

	// Type Struct

	type Address struct {
		Street   string
		Province string
	}

	type Student struct {
		Name    string
		Email   string
		Grade   string
		Weight  int
		Height  int
		Address Address
	}

	var student [3]Student
	student[0].Name = "Bambino"
	student[0].Email = "Bambino@Bambino"
	student[0].Grade = "A"
	student[0].Weight = 80
	student[0].Height = 175
	student[0].Address = Address{
		Street:   "Street1",
		Province: "Province1",
	}

	fmt.Println(student[0])

	studentMap := make(map[string]Student)
	studentMap["st01"] = Student{Name: "Bambino", Address: Address{
		Province: "Province1",
		Street:   "Street1",
	}}

	fmt.Println(studentMap)

	// Receiver argument
	ReceiverName()
	ReceiverSum()

	// Interface
	InterfaceSpeak()

	// Pointer
	x := 20
	x = 50

	fmt.Println(x)

	changeValue(&x)
	fmt.Println(x)

	var p *int = &x
	*p = 10
	fmt.Printf("x -> %d \n", x)
	fmt.Printf("p -> %d \n", *p)

	emp := Employee{Name: "Thanya", Salary: 50000}

	giveRaise(&emp, 5000)
	fmt.Println("emp", emp)

	// Error
	result, err := divideByZero(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

// Receiver argument

type Teacher struct {
	FirstName string
	LastName  string
}

func (s Teacher) FullName() string {
	return s.FirstName + " " + s.LastName
}

func ReceiverName() {
	teacher := Teacher{
		FirstName: "BamFirst",
		LastName:  "BamLast",
	}

	teacherFullName := teacher.FullName()

	fmt.Println(teacherFullName)
}

type IntSlice []int

func (s IntSlice) Sum() int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}

func ReceiverSum() {
	sum := IntSlice{1, 2, 3, 4, 5}.Sum()
	fmt.Println(sum)
}

// interface

type Speaker interface {
	Speak() string
	Walk() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello"
}
func (p Person) Walk() string {
	return "Hello Walk"
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Wolf!"
}
func (d Dog) Walk() string {
	return "Wolf! Walk"
}

func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func InterfaceSpeak() {
	dog := Dog{Name: "Buddy"}
	person := Person{Name: "Alice"}

	makeSound(dog)
	makeSound(person)
}

// Pointer
func changeValue(x *int) {
	*x = 30
}

type Employee struct {
	Name   string
	Salary int
}

func giveRaise(e *Employee, raise int) {
	e.Salary += raise
}

// Error

func divideByZero(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return a / b, nil
}

package data_type

import "fmt"

// Example struct
type Person struct {
	Name string
	Age  int
}

// Example interface
type Greeter interface {
	Greet() string
}

// Interface implementation
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func DataType() {
	// Const
	const pi float64 = 3.14
	fmt.Println("Pi:", pi)

	// String
	var name string = "Andi"
	fmt.Println("Name:", name)

	// Integer
	var age int = 25
	fmt.Println("Age:", age)

	// Unsigned integer
	var positiveNumber uint = 100
	fmt.Println("Positive Number:", positiveNumber)

	// Float
	var height float64 = 175.5
	fmt.Println("Height:", height)

	// Boolean
	var isStudent bool = true
	fmt.Println("Is Student:", isStudent)

	// Multiple variable declaration
	var a, b, c int = 1, 2, 3
	fmt.Println("a:", a, "b:", b, "c:", c)

	// Short declaration
	country := "Indonesia"
	fmt.Println("Country:", country)

	// Array
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)

	// Slice
	sliceNumbers := []int{10, 20, 30}
	sliceNumbers = append(sliceNumbers, 40)
	fmt.Println("Slice:", sliceNumbers)

	// Map
	personAge := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println("Map:", personAge)
	fmt.Println("Bob's age:", personAge["Bob"])

	// Struct
	person := Person{Name: "Dary", Age: 28}
	fmt.Println("Struct:", person)
	fmt.Println("Person Name:", person.Name)

	// Pointer
	ptr := &age
	fmt.Println("Pointer:", *ptr)
	*ptr = 26
	fmt.Println("Updated Age via Pointer:", age)

	// Interface
	var greeter Greeter = person
	fmt.Println(greeter.Greet())
}

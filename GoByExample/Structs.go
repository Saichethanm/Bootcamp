package main

import "fmt"

type Person struct {
	name string
	age  int
}

// constructor to the struct
func newPerson(name string, age int) *Person {
	newperson := Person{name: name, age: age}
	return &newperson
}

func main() {

	//Creating directly

	fmt.Println(Person{"a", 10})

	// If you wont mention columns its value is 0
	fmt.Println(Person{age: 26})

	//Using Constructor
	person1 := newPerson("Alice", 20)
	person2 := newPerson("Bob", 21)
	person3 := newPerson("Hero", 22)

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)

}

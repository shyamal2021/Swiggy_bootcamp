package main

import "fmt"

//There is no relationship between
//interface and struct here

type Animal interface {
	//method name
	//return type
	Speak() string

}

func (e Emp) Speak() string {
	return "Hello ....Hi..."
}

func (d Dog) Speak() string {
	return "Wof ....Hi..."
}

func (c Cat) Speak() string {
	return "Meow ....Hi..."
}

type Dog struct {}
type Emp struct {}
type Cat struct {}

func main()  {
	dog1 := Animal(Dog{})
	fmt.Println(dog1.Speak())

	animals := [] Animal{Emp{},Dog{},Cat{}}

	for _,animal := range animals {
		fmt.Println(animal.Speak())
	}

}

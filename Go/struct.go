package main

import "fmt"

type Dog struct {
	Breed string
	Weight int
}


func main(){
	dog1 := Dog{Breed: "Beagle",Weight: 5}
	fmt.Println(dog1)

	fmt.Printf("Breed : %v\nWeight : %v",dog1.Breed,dog1.Weight)


}

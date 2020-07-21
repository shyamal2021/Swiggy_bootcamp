package main

import "fmt"

func main(){
	var names [5] string

	names[0]="Shyam"
	names[1]="Aditya"
	names[2]="Abhi"
	names[3]="Anshu"
	names[4]="Shyamal"

	fmt.Println(names)

	var numbers = [3] int {2,3,4}
	fmt.Println(numbers,"\nLength is ",len(numbers))



}

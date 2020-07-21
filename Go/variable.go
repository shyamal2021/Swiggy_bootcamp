package main

import "fmt"

func main(){
	str1 := "This is string 1"
	str2 := "This is string 2 "
	str3 := "Hello there"
	aNumber := 54
	isTrue := true

	strLength, err := fmt.Println(str1, str2, str3)

	if err == nil {
		fmt.Println("The Length of the string is ", strLength)
	}

	fmt.Printf("Value of NUmber =  %v\n", aNumber)
	fmt.Printf("Value of isTrue =  %v\n", isTrue)
	fmt.Printf("Value of float NUmber =  %.2f\n", float64(aNumber) )
	fmt.Printf("Data Types =  %T , %T, %T\n", str1,aNumber,isTrue)
}
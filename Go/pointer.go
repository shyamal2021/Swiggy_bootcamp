package main

import "fmt"

func main(){
	var p1 *int

	v1 := 45
	p1 = &v1

	if p1 != nil {
		fmt.Println("Value of P", *p1)
	} else {
		fmt.Println("Sorry p is nil")
	}

	v2 := 50.65
	p2 := &v2
	fmt.Println("Value of p2",*p2)

}

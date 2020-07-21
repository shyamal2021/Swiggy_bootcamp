package main

import (
	"fmt"
	"math"
)

func main(){

	i1, i2, i3 := 10, 20, 30
	intSum := i1+i2+i3

	fmt.Println(intSum)

	f1, f2, f3 := 23.12, 32.32, 23.432
	floatSum := f1+f2+f3

	fmt.Println(floatSum)

	radius := 14.5
	circumference := radius * math.Pi

	fmt.Println("The Circumference is = ",circumference)



}

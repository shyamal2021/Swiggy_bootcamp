package main

import (
	"fmt"
	"sort"
)

func main(){
	var colors = [] string {"Red","blue","Green"}
	fmt.Println(colors)

	colors = append(colors,"Yellow")
	fmt.Println(colors)

	colors = colors[1:len(colors)-1]
	fmt.Println(colors)

	ints := make([]int,5,5) //parmeter will initial capacity and incresing capacity
	ints[0] = 23
	ints[1] = 78
	ints[2] = 28
	ints[3] = 27
	ints[4] = 26
	fmt.Println("After Adding 5 elements the capacity will be ", cap(ints))
	fmt.Println(ints)

	ints = append(ints,100)
	fmt.Println("After Adding 6th elements the capacity will be ", cap(ints))
	fmt.Println(ints)

	//Ascending order
	sort.Ints(ints)
	fmt.Println(ints)

	//one way to descend order
	sort.Slice(ints, func(i, j int) bool {
						return ints[i] > ints[j]
					  })
	fmt.Println(ints)

	//Descending order
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints)


}
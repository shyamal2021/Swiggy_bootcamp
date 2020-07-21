package main

import (
	"fmt"
	"time"
)

func main()  {
	t := time.Date(2020,time.July,20,10,10,10,10,time.UTC)
	fmt.Printf("Time is %s\n ",t)

	now := time.Now()
	fmt.Printf("Time Now is %s\n ",now)

	//day
	fmt.Println("Day is ",t.Day())

	//Month
	fmt.Println("Month is ",t.Month())

	//Weekday
	fmt.Println("Week Day ", t.Weekday())

	tomorrow := t.AddDate(0,0,1)
	fmt.Printf("Tommorow is %s\n",tomorrow)

	fmt.Printf("Tomorrow Date : %v,Month : %v\n",tomorrow.Day(),tomorrow.Month())

	//Wednesday ,Jul 22, 2020
	fmt.Println(tomorrow.Format(time.RFC850))
	fmt.Println(tomorrow.Format(time.RFC1123))

	//7/22/2020
	fmt.Println(tomorrow.Format("1/2/09"))
	fmt.Println(tomorrow.Format("Monday, Jan"))









}

package main

import (
	"bufio"
	"fmt"
	"os"
)

//1. Write a program which has the
//struct Emp -> empid, empname, salary, address
//type Address struct { hno, street, city, pin }
//Accept the values from the end-user, for initial 5 employees, you should give an option for use to quit in between Q (if Condition)
//then display the details
//
//There is a need for contracting to eat, sleep, breath -> interface called Animal
//If the user enters more than 5 employees then the size of the slice should be +5

type Animal interface {
	eat() string
	sleep() int
	breath() string
}

func (e Emp) eat() string {
	return "Rice,fruits etc"
}

func (e Emp) sleep() int {
	return 8
}

func (e Emp) breath() string {
	return "Nasal Respiratory System"
}


type Emp struct {
	empid string
	empname string
	salary int
	address *Address
}

type Address struct {
	hno string
	street string
	city string
	pin int
}


func main(){
	reader := bufio.NewReader(os.Stdin)
	emps := []*Emp{}
	i:=0

	for {
		emp := new(Emp)

			fmt.Println("Enter Empid")
			emp.empid ,_ = reader.ReadString('\n')

			fmt.Println("Enter Employee Name")
			emp.empname ,_ = reader.ReadString('\n')

			fmt.Println("Enter Employee Salary")
			fmt.Scan(&emp.salary)

		adr := new(Address)

			fmt.Println("Enter house number")
			adr.hno ,_ = reader.ReadString('\n')

			fmt.Println("Enter Street name")
			adr.street ,_ = reader.ReadString('\n')

			fmt.Println("Enter Employee City Name")
			adr.city ,_ = reader.ReadString('\n')

			fmt.Println("Enter pin number")
			//adr.pin ,_ = reader.ReadString('\n')
			fmt.Scan(&adr.pin)

		emp.address = adr

		emps = append(emps,emp)

		fmt.Println("To enter more emplyees : (0/1) ")
		response := 0
		fmt.Scan(&response)
		if response == 0 {
			break
		}
		i++
	}

	for i=0;i<len(emps);i++ {
		fmt.Println("Empid = ",emps[i].empid,"Empname = ",emps[i].empname,"Salary = ",emps[i].salary)
		fmt.Println("House Number = ",emps[i].address.hno,"Street name = ",emps[i].address.street,"City Name = ",emps[i].address.city)
		fmt.Println("Pin Number = ",emps[i].address.pin)
		fmt.Println("Eating Habit : ",emps[i].eat())
		fmt.Println("Avg sleeping Duration : ",emps[i].sleep())
		fmt.Println("Breathing Process : ",emps[i].breath())
	}

}

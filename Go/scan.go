package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter ur salary")
	strSalary , _ := reader.ReadString('\n')

	fmt.Println("Enter Months")
	strMonth , _ := reader.ReadString('\n')

	salary, _ := strconv.ParseFloat(strings.TrimSpace(strSalary),64)
	month, _ := strconv.ParseInt(strings.TrimSpace(strMonth),10 ,64)

	TotalSalary := salary * float64(month)

	fmt.Println("Total Salary is = ", TotalSalary)
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter Employee Type")
	fmt.Println("1. Full Time")
	fmt.Println("2. Contractor")
	fmt.Println("3. Freelancer")
	var s, t int
	fmt.Scan(&s)
	switch s {
	case 1:
		fmt.Println("Enter no. of Working days")
		fmt.Scan(&t)
		fmt.Printf("Salary is %v", t*500)
	case 2:
		fmt.Println("Enter no. of Working days")
		fmt.Scan(&t)
		fmt.Printf("Salary is %v", t*100)
	case 3:
		fmt.Println("Enter no. of Working hours")
		fmt.Scan(&t)
		fmt.Printf("Salary is %v", t*20)
	default:
		fmt.Println("Invalid Input")
	}

}

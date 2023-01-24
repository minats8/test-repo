package main

import (
	art "dummy/article"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid no of arguments...")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "s":
		ShowData(art.Student{
			Name:   "Sidhant",
			Course: "MCA",
			Rollno: 1,
		})
	case "e":
		ShowData(art.Employee{
			Name:   "Reetwik",
			UserId: 2,
		})
	default:
		fmt.Println("Invalid argument...")
		os.Exit(1)
	}
}

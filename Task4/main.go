package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Invalid Number of Arguments...")
		os.Exit(1)
	}
	var cmdArg string
	if len(os.Args) == 1 {
		cmdArg = ""
	} else {
		cmdArg = os.Args[1]
	}
	switch cmdArg {
	case "d":
		fmt.Println("Proceeding with database...")
		Menu(cmdArg)
	case "":
		fmt.Println("Proceeding with File System...")
		Menu(cmdArg)
	default:
		fmt.Println("Invalid Argument...")
		os.Exit(1)
	}
}

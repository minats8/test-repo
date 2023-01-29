package main

import (
	"fmt"
	"os"
	structure "task4/dataStructure"
	"task4/models"
)

func Menu(cmdArg string) {
	fmt.Println("Select Option ...\n1.View Storage\n2.Insert into Storage\n3.Exit\nEnter Your Choice: ")

	var choice int
	fmt.Scan(&choice)

	var fs structure.FileStructure
	var db structure.DatabaseStructure
	var err error

	switch choice {
	case 1:
		if cmdArg == "" {
			err = Fetch(fs)
			if err != nil {
				os.Exit(1)
			}
		} else {
			err = Fetch(db)
			if err != nil {
				os.Exit(1)
			}
		}
	case 2:
		var postNo string
		fmt.Println("Enter the Post Number you want to Insert: ")
		fmt.Scan(&postNo)
		var data models.Data
		data, err = LoadPageData(postNo)
		if err != nil {
			os.Exit(1)
		}
		if cmdArg == "" {
			err = WrtieOnStorage(fs, data)
			if err != nil {
				os.Exit(1)
			}
			fmt.Println("Data Written on file successfully...")
		} else {
			err = WrtieOnStorage(db, data)
			if err != nil {
				os.Exit(1)
			}
			fmt.Println("Data Inserted into database successfully...")
		}
	case 3:
		fmt.Println("Thank You for visit...")
		os.Exit(0)
	default:
		fmt.Println("Invalid Option Selected...")
		os.Exit(1)
	}
}

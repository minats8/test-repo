package main

import (
	art "dummy/article"
	"fmt"
)

func ShowData(person art.Person) {
	// fmt.Println(art.Person.PutData(person))
	fmt.Println(person.PutData())
}

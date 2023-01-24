package art

import (
	"strconv"
)

type Employee struct {
	Name   string
	UserId int
}

type Student struct {
	Name   string
	Course string
	Rollno int
}

type Person interface {
	PutData() string
}

func (s Student) PutData() string {
	data := "Name: " + s.Name + " Course: " + s.Course + " Rollno: " + strconv.Itoa(s.Rollno)
	return data
}

func (e Employee) PutData() string {
	data := "Name: " + e.Name + " UserID: " + strconv.Itoa(e.UserId)
	return data
}

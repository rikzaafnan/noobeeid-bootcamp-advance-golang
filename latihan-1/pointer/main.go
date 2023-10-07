package main

import "fmt"

func main() {
	dataStudent := Student{
		Name:  "carly",
		Class: 1,
	}

	dataStudent.SetMyName("brandong")
	dataStudent.CallMyName()
	dataStudent.SetMyName("brandong 2")
	dataStudent.CallMyName()

}

type Student struct {
	Name  string
	Class int
}

type StudentInterface interface {
	SetMyName(name string)
	CallMyName()
}

func (s *Student) SetMyName(name string) {
	s.Name = name
}

func (s Student) CallMyName() {

	fmt.Println("Hello, My Name is", s.Name)

}

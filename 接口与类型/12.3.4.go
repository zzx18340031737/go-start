package main

import "fmt"

type IPerson interface {
	Speak()
}

type IStudent interface {
	IPerson
	Study()
}

type ITeacher interface {
	IPerson
	Teach()
}

type Student struct {
	Name string
}

func (s *Student)Speak(){
	fmt.Println("My name is",s.Name)
}

func (s *Student)Study(){
	fmt.Println(s.Name,"is studying")
}

type Teacher struct {
	Name string
}

func (t *Teacher)Speak(){
	fmt.Println("My name is",t.Name)
}

func (t *Teacher)Teach(){
	fmt.Println(t.Name,"is teaching")
}

func main(){
	var stu Student = Student{"Xiao Ming"}
	var teacher Teacher = Teacher{"Mr. Li"}

	stu.Speak()
	stu.Study()
	teacher.Speak()
	teacher.Teach()
}
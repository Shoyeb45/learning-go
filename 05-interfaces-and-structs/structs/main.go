package main

import "fmt"

// we define struct and interfaces using type


type student struct {
	name string
	rollNumber string
	studentCourse course 
}

type course struct {
	code int32
	name string
}

// function associated with the struct

func (st student) getName() string {
	return  st.name;
}

func main() {
	st1 := student{name: "Shoyeb", rollNumber: "123", studentCourse: course{12, "DA"}};

	// we can change the value like:

	st1.rollNumber = "3212";

	fmt.Println(st1);

	fmt.Println(st1.getName());
}
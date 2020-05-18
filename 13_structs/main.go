//A struct is a collection of fields.
package main

import "fmt"

type student struct {
	name   string
	rollNo int
}

// This are struct literal which denotes a newly allocated struct value by listing the values of its fields.
var (
	s1 = student{"Harry", 1}  // has type student
	s2 = student{name: "Ron"} // rollNo:0 is implicit
	s3 = student{}            // name: "", rollNo: 0 is implicit
	s4 = &student{"Rick", 2}  // has type *student
)

//Go supports methods defined on struct types.
func (s *student) updateStudent(name string, rollNo int) {
	s.name = name
	s.rollNo = rollNo
}

func (s student) getStudentInfo() (string, int) {
	return s.name, s.rollNo
}

func main() {
	//Print struct literals
	fmt.Println(s1, s2, s3, s4)

	fmt.Println()

	// accessing fields
	s := student{"Jessica Rich", 15}
	fmt.Println(s)
	s.name = "Jessica"
	fmt.Println(s)

	fmt.Println()

	//Pointers to struct
	p := &s
	fmt.Println("Name: ", (*p).name)   //access field through *p
	fmt.Println("Roll NO: ", p.rollNo) //access field through p

	fmt.Println()

	//methods on struct
	s.updateStudent("Richie", 101)
	fmt.Println(s)

	name, rollNo := s.getStudentInfo()
	fmt.Println(name, rollNo)
}

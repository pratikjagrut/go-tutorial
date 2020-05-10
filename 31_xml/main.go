//Go offers built-in support for XML and XML-like formats with the encoding.xml package.
package main

import (
	"encoding/xml"
	"fmt"
)

type student struct {
	XMLName xml.Name `xml:"student"`
	Name    string   `xml:"name"`
	RollNo  int      `xml:"rollNo"`
	Subject []string `xml:"subject"`
}

func main() {
	student1 := &student{
		Name:    "Billy",
		RollNo:  10,
		Subject: []string{"math", "comp"},
	}

	out1, _ := xml.Marshal(student1) //Convert struct to xml
	fmt.Println(string(out1))
	fmt.Println()

	out, _ := xml.MarshalIndent(student1, " ", " ") //Convert struct to xml and indent it
	fmt.Println(string(out))
	fmt.Println()

	fmt.Println(xml.Header + string(out)) //Add generic header
	fmt.Println()

	var s student

	//Use Unmarhshal to parse a stream of bytes with XML into a data structure.
	if err := xml.Unmarshal(out, &s); err != nil {
		panic(err)
	}
	fmt.Println(s)
	fmt.Println()

	student2 := &student{
		Name:    "Jessica",
		RollNo:  2,
		Subject: []string{"Geology", "Astronomy"},
	}

	type Nesting struct {
		XMLName  xml.Name   `xml:"nesting"`
		Students []*student `xml:"class>section>student"` //this tells encoder to nest all students under <class><section>
	}

	nesting := &Nesting{
		Students: []*student{student1, student2},
	}

	out, _ = xml.MarshalIndent(nesting, " ", " ") //Convert struct to xml and indent it
	fmt.Println(string(out))
}

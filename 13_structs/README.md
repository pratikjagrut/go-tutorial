# Structure

In simple words, type struct is a collection of fields, similar to a class in object oriented programming.

We use `struct` keyword to create a new structure type.

```go
type structName struct{
    fieldName1 fieldType1
    fieldName2 fieldType2
    fieldName3 fieldType3
    .
    .
}
```

The zero value of a struct is a struct with fields of its own zero values.

## Creating a struct

We'll create type student of type struct here

```go
type student struct {
    name   string
    rollNo int
}
```

## Accessing fields

Create struct objects through which we can access struct fields.

Struct fields are accessed using a dot.

```go
s1 := student{"Harry", 1}
fmt.Println(s1.name) //This will print name
fmt.Println(s1.rollNo) //This will print roll no
```

## Pointers to struct

```go
ptr := &s1
```

After assigning a pointer to struct it's field can be accessed using `(*ptr).name` but it seems a little bit unhandy so GO allow us to just use `ptr.name`.

```go
fmt.Println((*ptr).name)
fmt.Println(ptr.rollNo)
```

## Struct literal

```go
s1 := student{"Harry", 1}  // has type student
s2 := student{name: "Ron"} // rollNo:0 is implicit
s3 := student{}            // name: "", rollNo: 0 is implicit
s4 := &student{"Rick", 2}  // has type *student
```

## Methods on struct

GO supports defining method on the struct.

```go
func (s *student) updateStudent(name string, rollNo int) {
    s.name = name
    s.rollNo = rollNo
}
```

Now, this method can only be accessed by struct objects.

```go
s1.updateStudent("Richie", 101)
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
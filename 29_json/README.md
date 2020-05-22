# JSON

JSON (JavaScript Object Notation) is a simple data interchange format. Syntactically it resembles the objects and lists of JavaScript. It is most commonly used for communication between web back-ends and JavaScript programs running in the browser, but it is used in many other places, too. Its home page, [json.org](https://www.json.org/json-en.html), provides a wonderfully clear and concise definition of the standard.

This is an example of a simple JSON.

```
{
    "firstName": "John",
    "lastName": "Wick",
    "nickname" : "Baba Yaga/The Boogeyman",
    "age": "He has no age",
    "height": 6.1,
    "isMale": true,
    "profession": "Hitman (formerly) Vigilante",
    "fightingStyle": [  
                        "Japanese Jiu-jitsu",
                        "Brazilian jiu-jitsu",
                        "Judo",
                        "Sambo"
                     ]
    "languages": [
                    "English",
                    "Russian",
                    "Italian",
                    "sign language"
                 ]
}
```

With the JSON package, it's easy to read and write JSON data from your Go programs.
Package JSON implements encoding and decoding of JSON as defined in RFC 7159.

```go
import "encoding/json"
```
## Encoding 

To encode JSON data we use the Marshal function.

This is a signature of Marshal function.

`func Marshal(v interface{}) ([]byte, error)`

It generates array []byte containing JSON data.

e.g.

```go
str, _ := json.Marshal("Golang")
fmt.Println(str) 
```

Output: 
```
[34 71 111 108 97 110 103 34]
```

Converting it to string

```go
fmt.Println(string(str))
```

Output: 
```
"Golang"
```
GO struct with data

```go
type student struct {
    //Only exported fields will be encoded/decoded in JSON.
    //Fields must start with capital letters to be exported.
    Name   string
    RollNo int `json:"RegNo"` //Json msg will contain RegNo, instead of RollNo
}
```

```go
studentB, _ := json.Marshal(&student{"Jon", 1})
fmt.Println(string(studentB))
```

Output: 
```
{"Name":"Jon","RegNo":1}
```

## Unmarshal

Unmarshalling above student JSON.

```go
stud1 := student{}
json.Unmarshal(studentB, &stud1)
fmt.Println(stud1)
```

Output:

```
{Jon 1}
```

 ***You can refer main.go file for examples***

To run:
```
go run main.go
```
# Read from stdin


***Create new reader.***

NewReader returns a new Reader whose buffer has the default size

```go
reader := bufio.NewReader(os.Stdin)
```

***Read using ReadString.***

Signature: `func (*bufio.Reader).ReadString(delim byte) (string, error)`

`ReadString` reads until the first occurrence of `delim in the input`, returning a string containing the data up to and including the delimiter. If ReadString encounters an error before finding a delimiter, it returns the data read before the error and the error itself (often io.EOF). ReadString returns err != nil if and only if the returned data does not end in delim. For simple uses, a Scanner may be more convenient.

```go
fmt.Println("First name: ")
firstName, _ := reader.ReadString('\n') // Reads unitl \n char is found
```

***Read using Scanner.***

```go
scanner := bufio.NewScanner(os.Stdin)
var lastName string
fmt.Println("Last name: ")
for scanner.Scan() {
    lastName = scanner.Text()
    break
}
```

***You can refer main.go file for examples***

To run:
```
go run main.go
First name: 
James
Last name: 
Bond
Hello, James Bond
```
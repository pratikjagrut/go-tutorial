# Writing file

In this tutorial, we will learn how to write data to files using Go.

## Writing string to a file

One of the most common files writing operations is writing a string to a file.

***First create file with `os.Create(arg)`***

`Create` creates or truncates the named file. If the file already exists, it is truncated. If the file does not exist, it is created with `mode 0666` (before umask). If successful, methods on the returned File can be used for I/O; the associated `file descriptor has mode O_RDWR`. If there is an error, it will be of type `*PathError`.

```go
f, err := os.Create("string.txt")
if err != nil {
    panic(err)
}
```

***Close file***

It important to close the file after use. `defer` will invoke this call at the end.

```go
defer f.Close()
```

***Write string to file***

`WriteString` is like Write, but writes the contents of string s rather than a slice of bytes. Returns number of byte written

```go
n, err := f.WriteString("Go is fun!")
if err != nil {
    panic(err)
}
```

## Writing bytes to a file

Writing bytes to a file is quite similar to writing a string to a file. We will use the Write method to write bytes to a file. 

***Create a file and deferred Close.***

```go
f2, err := os.Create("byte.txt")
if err != nil {
    panic(err)
}

defer f2.Close()
```

***Write byte slice to the byte.txt***

Write a byte using Write function.

Signature of Write fucntion: 

`func (*os.File).Write(b []byte) (n int, err error)`

`Write` writes `len(b)` bytes to the File. It returns the number of bytes written and an error if any. Write returns a non-nil error when n != len(b).

```go
data := []byte{71, 111, 32, 105, 115, 32, 102, 117, 110, 33}
n2, err := f2.Write(data)
if err != nil {
    panic(err)
}
```

## Writing file line by line

Another common file operation is the need to write strings to a file line by line.

***Create a file and deferred Close.***

```go
f3, err := os.Create("lineByLine.txt")
if err != nil {
    panic(err)
}
defer f3.Close()
```

***Write text line by line to file.***

Text: 

```go
text := []string{
    "Go is a statically typed,",
    "compiled programming language designed at Google",
    "by Robert Griesemer, Rob Pike, and Ken Thompson.",
}
```

Write lines by using fmt.Fprintln function.

Signature : 

`func fmt.Fprintln(w io.Writer, a ...interface{}) (n int, err error)`

`Fprintln` formats using the default formats for its operands and writes to `w`. Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.

```go
n3 := 0 // Number of bytes
for _, str := range text {
    n3, err = fmt.Fprintln(f3, str)
    if err != nil {
        panic(err)
    }
}
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
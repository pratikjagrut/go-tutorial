# Reading Files

File reading is one of the most common operations performed in any programming language. In this tutorial, we will learn about how files can be read using Go.

## Reading an entire file into memory

One of the most basic file operations is reading an entire file into memory. This is done with the help of the ReadFile function of the `ioutil` package.

```go
data, err := ioutil.ReadFile("data.txt")
if err != nil {
    panic(err)
}
fmt.Println(string(data))
```

Output:

```
This is GOlang!
```

## Reading a file with more control

We don't always want to read the entire file in memory especially when you've low memory, sometimes you need more control over how much to read or from where to start reading.

**Open file**

`os` package provides `open` file function.

`Open` opens the named file for reading. If successful, methods on the returned file can be used for reading; the associated file descriptor has mode O_RDONLY. If there is an error, it will be of type *PathError.

```go
f, err := os.Open("data.txt")
if err != nil {
    panic(err)
}
```

Now here `f` is a pointer to  `os.File`.

**Close file**

It is good practice to close the file as soon as we finish operations on it.

`Close` closes the File, rendering it unusable for I/O. On files that support SetDeadline, any pending I/O operations will be cancelled and return immediately with an error. Close will return an error if it has already been called.

```go
defer f.Close()
```

**Restricting number of bytes to read from the file using byte slice**

Allow 20 bytes to be read from the file, our entire file is of 15 bytes

```go
b1 := make([]byte, 20)
n1, err := f.Read(b1)
if err != nil {
    panic(err)
}
fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
```

`n1` is an actual number of bytes read, `n1<20`, so it will read the whole file.

Output:

```
15 bytes: This is GOlang!
```

**Seek read pointer to known location**

We can control from which byte we want to read the file, i.e we can also Seek to a known location in the file and read from there. 

`Seek(x ,y)`, `x = offset` for the next read,

Seek sets the offset for the next Read or Write on file to offset, interpreted according to `whence: 0 means relative to the origin of the file, 1 means relative to the current offset, and 2 means relative to the end`. It returns the new offset and an error if any. The behaviour of Seek on a file opened with O_APPEND is not specified.

```go
out, err := f.Seek(5, 0)
if err != nil {
    panic(err)
}
b2 := make([]byte, 2) //Allow 2 bytes
n2, err := f.Read(b2)
if err != nil {
    panic(err)
}
fmt.Printf("read ptr @ %d:  \t%v\n", out, string(b2[:n2]))
```

Output:

```
read ptr @ 5:   is
```

**ReadAtLeast**

```go
func io.ReadAtLeast(r io.Reader, buf []byte, min int) (n int, err error)
```

`ReadAtLeast` reads from r into a buffer until it has read at least min bytes. It returns the number of bytes copied and an error if fewer bytes were read. The error is EOF only if no bytes were read. If an EOF happens after reading fewer than min bytes, ReadAtLeast returns ErrUnexpectedEOF. If min is greater than the length of the buffer, ReadAtLeast returns ErrShortBuffer. On return, n >= min if and only if err == nil. If r returns an error having read at least min bytes, the error is dropped.

```go
_, err = f.Seek(8, 0) //seeks read pointer to the 8 byte
if err != nil {
    panic(err)
}
b3 := make([]byte, 7) //Allow 7 bytes to be read
n3, err := io.ReadAtLeast(f, b3, 2) // ReadAtLeast reads from f into buffer b3 until it has read at least 2 bytes.
if err != nil {
    panic(err)
}
fmt.Printf("%d bytes read: \t%s\n", n3, string(b3))
```

Output:

```
7 bytes read:   GOlang!
```

##  Reading a file in small chunks

When the size of the file is extremely large it doesn't make sense to read the entire file into memory especially if you are running low on RAM. A more optimal way is to read the file in small chunks. This can be done with the help of the `bufio package`.

*Create a new buffered reader*

```go
r4 := bufio.NewReader(f)
```

*byte slice of length and capacity 2 into which the bytes of the file will be read.*

```go
b4 := make([]byte, 2)
```

The `Read method` reads up to `len(b4)` bytes i.e up to 2 bytes and returns the number of bytes read. We store the bytes returned in a variable `n`. The slice is read from index 0 to n-1, i.e up to the number of bytes returned by the Read method and printed.

Once the end of the file is reached, it will return an EOF error.

```go
for {
    n, err := r4.Read(b4)
    if err != nil {
        fmt.Println("Error reading file:", err)
        break
    }
    fmt.Println(string(b4[0:n]))
}
```

Output: 

```

Reading a file in a chunk of 2
Th
is
 i
s 
GO
la
ng
!
Error reading file: EOF
```

## Reading a file line by line

In the section, we will discuss how to read a file line by line using Go. This can be done using the bufio package.

Open a file and deferred close.

```go
f1, err := os.Open("file.txt")
if err != nil {
    panic(err)
}

defer f1.Close()
```

Create a new scanner from the file

```go
s := bufio.NewScanner(f1)
```

Scan the file and read it line by line.

```go
for s.Scan() {
    fmt.Println(s.Text())
}
err = s.Err()
if err != nil {
    panic(err)
}
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
# Command-line arguments

Command-line arguments are a common way to parameterize the execution of programs.

`os.Args` provides access to raw command-line arguments. 

```go
args := os.Args
```

Note that the `first value in this slice is the path to the program`, and os.Args[1:] holds the arguments to the program.

```go
fmt.Println("Path of the program: ", args[0]) // First argument is always path of program
```

Output:

```
go run main.go

Path of the program:  /tmp/go-build539456514/b001/exe/main
```

Arguments from `os.Args[1:]`.

```go
for i, arg := range args {
    fmt.Printf("arg index: %d, arg: %v\n", i, arg)
}
```

Output:

```
go run main.go a b c d

arg index: 0, arg: /tmp/go-build426055485/b001/exe/main
arg index: 1, arg: a
arg index: 2, arg: b
arg index: 3, arg: c
arg index: 4, arg: d
```